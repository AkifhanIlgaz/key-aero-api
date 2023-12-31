package cfg

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost     string `mapstructure:"PSQL_HOST"`
	PostgresPort     string `mapstructure:"PSQL_PORT"`
	PostgresUser     string `mapstructure:"PSQL_USER"`
	PostgresPassword string `mapstructure:"PSQL_PASSWORD"`
	PostgresDBName   string `mapstructure:"PSQL_DBNAME"`
	PostgresSSLMode  string `mapstructure:"PSQL_SSLMODE"`

	RedisUrl string `mapstructure:"REDIS_URL"`

	Port string `mapstructure:"PORT"`

	AccessTokenPrivateKey string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey  string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge     int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`

	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUsername string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	return &config, nil
}
