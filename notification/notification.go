package notification

import (
	"fmt"

	"github.com/kmdkuk/gote/notification/slack"
	"github.com/kmdkuk/gote/notification/twitter"
)

func Notification(dest, message string) error {
	if dest == "twitter" {
		return twitter.Post(message)
	} else if dest == "slack" {
		return slack.Post(message)
	} else {
		return fmt.Errorf("invalid notification destination")
	}
}
