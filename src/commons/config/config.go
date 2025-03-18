package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
}

type AppConfig struct {
	Host  string
	Port  int
	Level string
}

type DBConfig struct {
	Port     int
	Host     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	AccessKey  SecretKey
	RefreshKey SecretKey
}

type SecretKey struct {
	Key               string
	ExpireTimeInHours int
}

func New(config *viper.Viper, level string) *Config {
	if level == "local" {
		config.SetConfigName("local")
		config.SetConfigType("env")
		config.AddConfigPath("env/.")

		err := config.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	} else {
		config.AutomaticEnv()
	}

	return &Config{
		App: AppConfig{
			Host:  config.GetString("APP_HOST"),
			Port:  config.GetInt("APP_PORT"),
			Level: config.GetString("APP_LEVEL"),
		},
		DB: DBConfig{
			Port:     config.GetInt("POSTGRES_PORT"),
			Host:     config.GetString("POSTGRES_HOST"),
			User:     config.GetString("POSTGRES_USER"),
			Password: config.GetString("POSTGRES_PASSWORD"),
			Name:     config.GetString("POSTGRES_DB"),
		},
		JWT: JWTConfig{
			AccessKey: SecretKey{
				Key:               config.GetString("JWT_ACCESS_KEY"),
				ExpireTimeInHours: config.GetInt("JWT_ACCESS_TIME"),
			},
			RefreshKey: SecretKey{
				Key:               config.GetString("JWT_REFRESH_KEY"),
				ExpireTimeInHours: config.GetInt("JWT_REFRESH_TIME"),
			},
		},
	}
}
