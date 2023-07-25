package config

import (
	"os"

	"github.com/spf13/viper"
)

const (
	SMTP_HOST          = "smtp.gmail.com"
	SMTP_PORT          = 587
	SMTP_SENDER_NAME   = "Quliku <no-reply@qulikuindonesia>"
	SMTP_AUTH_EMAIL    = "qulikuindonesia@gmail.com"
	SMTP_AUTH_PASSWORD = "jxtdtpeebqoeqvzg"
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

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		var config EmailConfig
		if err := viper.Unmarshal(&config); err != nil {
			return nil, err
		}

		return &config, nil
	} else {
		return &EmailConfig{
			Host:         SMTP_HOST,
			Port:         SMTP_PORT,
			SenderName:   SMTP_SENDER_NAME,
			AuthEmail:    SMTP_AUTH_EMAIL,
			AuthPassword: SMTP_AUTH_PASSWORD,
		}, nil
	}

	// return &EmailConfig{}, nil
}
