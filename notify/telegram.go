package notify

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/codex416/sub-audit-gateway/config"
)


func SendTelegram(cfg *config.Config, message string) error {


	if !cfg.Telegram.Enable {

		return nil

	}



	api := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		cfg.Telegram.BotToken,
	)



	data := url.Values{}

	data.Set(
		"chat_id",
		cfg.Telegram.ChatID,
	)

	data.Set(
		"text",
		message,
	)



	_, err := http.PostForm(
		api,
		data,
	)


	return err

}
