package services

import (
	"context"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
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

}

func (service *TokenService) ParseAccessToken(token string) (*jwt.StandardClaims, error) {

}

func (service *TokenService) GenerateRefreshToken(uid string) (string, error) {

}

func (service *TokenService) DeleteRefreshToken(refreshToken string) error {

}
