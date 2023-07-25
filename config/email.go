package config

import (
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
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()

	var emailConfig EmailConfig
	if err := viper.Unmarshal(&emailConfig); err != nil {
		return nil, err
	}

	return &emailConfig, nil
}
