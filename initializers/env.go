package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// database
	DBHost string `mapstructure:"POSTGRES_HOST"`
	DBUser string `mapstructure:"POSTGRES_USER"`
	DBPass string `mapstructure:"POSTGRES_PASSWORD"`
	DBName string `mapstructure:"POSTGRES_DB"`
	DBPort string `mapstructure:"POSTGRES_PORT"`
	
	// app
	ServerPort string `mapstructure:"APP_PORT"`

	// JWT Access Token
	AccessTokenPrivate string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublic string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenExpiresIn time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`

	// JWT Refresh Token 
	RefreshTokenPrivate string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublic string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() 
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
