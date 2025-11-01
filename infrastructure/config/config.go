package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                 string
	DatabaseURL          string
	JwtSecret            string
	TelegramBotToken     string
	TelegramChatID       string
	DefaultCheckInterval int
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{
		Port:                 viper.GetString("PORT"),
		DatabaseURL:          viper.GetString("DATABASE_URL"),
		JwtSecret:            viper.GetString("JWT_SECRET"),
		TelegramBotToken:     viper.GetString("TELEGRAM_BOT_TOKEN"),
		TelegramChatID:       viper.GetString("TELEGRAM_CHAT_ID"),
		DefaultCheckInterval: viper.GetInt("DEFAULT_CHECK_INTERVAL"),
	}

	return c, nil
}
