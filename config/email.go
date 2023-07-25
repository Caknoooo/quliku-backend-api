package config

import (
	"os"
	"github.com/spf13/viper"
)

type EmailConfig struct {
	Host         string `mapstructure:"SMTP_HOST"`
	Port         int    `mapstructure:"SMTP_PORT"`
	SenderName   string `mapstructure:"SMTP_SENDER_NAME"`
	AuthEmail    string `mapstructure:"SMTP_AUTH_EMAIL"`
	AuthPassword string `mapstructure:"SMTP_AUTH_PASSWORD"`
}

func NewEmailConfig() (*EmailConfig, error) {
	if os.Getenv("APP_ENV") != "Production" {
		viper.SetConfigFile(".env")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var emailConfig EmailConfig
	if err := viper.Unmarshal(&emailConfig); err != nil {
		return nil, err
	}

	return &emailConfig, nil
}
