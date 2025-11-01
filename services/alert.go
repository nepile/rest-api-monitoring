package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nepile/api-monitoring/config"
)

func SendTelegramAlert(cfg *config.Config, message string) error {
	if cfg.TelegramBotToken == "" || cfg.TelegramChatID == "" {
		return nil
	}

	api := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramBotToken)
	data := url.Values{}
	data.Set("chat_id", cfg.TelegramChatID)
	data.Set("text", message)
	data.Set("parse_mode", "HTML")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Post(api, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}
