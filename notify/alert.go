package notify

import (
	"fmt"

	"github.com/codex416/sub-audit-gateway/config"
)


func SendAlert(
	cfg *config.Config,
	title string,
	content string,
) error {


	message := fmt.Sprintf(
		"🚨 %s\n\n%s",
		title,
		content,
	)



	return SendTelegram(
		cfg,
		message,
	)

}
