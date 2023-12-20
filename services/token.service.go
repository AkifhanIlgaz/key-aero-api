package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/thanhpk/randstr"
)

type TokenService struct {
	redisClient *redis.Client
	ctx         context.Context
	config      *cfg.Config
}

func NewTokenService(ctx context.Context, config *cfg.Config, redisClient *redis.Client) *TokenService {
	return &TokenService{
		redisClient: redisClient,
		ctx:         ctx,
		config:      config,
	}
}

func (service *TokenService) GenerateAccessToken(uid string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(service.config.AccessTokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	now := time.Now().UTC()

	claims := jwt.StandardClaims{}
	claims.Subject = uid
	claims.ExpiresAt = now.Add(service.config.AccessTokenExpiresIn).Unix()
	claims.IssuedAt = now.Unix()
	claims.NotBefore = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	return token, nil
}

func (service *TokenService) ParseAccessToken(token string) (*jwt.StandardClaims, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(service.config.AccessTokenPublicKey)
	if err != nil {
		return nil, fmt.Errorf("parse access token: %w", err)
	}

	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return jwt.StandardClaims{}, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})
	if err != nil {
		return nil, fmt.Errorf("parse access token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims, nil
}

func (service *TokenService) GenerateRefreshToken(uid string) (string, error) {
	refreshToken := randstr.String(32)

	err := service.redisClient.Set(service.ctx, refreshToken, uid, 7*time.Hour*24).Err()
	if err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	return refreshToken, nil
}

func (service *TokenService) DeleteRefreshToken(refreshToken string) error {
	if err := service.redisClient.Del(service.ctx, refreshToken).Err(); err != nil {
		return fmt.Errorf("delete refresh token: %w", err)
	}

	return nil
}
