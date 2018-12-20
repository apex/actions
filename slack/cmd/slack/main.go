package main

import (
	"fmt"
	"log"
	"os"

	"github.com/apex/actions/slack"
)

func main() {
	var msg slack.Message

	// read message
	err := slack.ReadMessage("slack.json", &msg)

	if os.IsNotExist(err) {
		log.Fatalf("Missing ./slack.json file, a previous action should populate it.")
	}

	if err != nil {
		log.Fatalf("error reading message: %s", err)
	}

	// webhook
	webhook := os.Getenv("SLACK_WEBHOOK_URL")

	if webhook == "" {
		log.Fatalf("Missing SLACK_WEBHOOK_URL environment variable")
	}

	// channel
	if s := os.Getenv("SLACK_CHANNEL"); s != "" {
		msg.Channel = s
	}

	// username
	if s := os.Getenv("SLACK_USERNAME"); s != "" {
		msg.Username = s
	}

	// icon
	if s := os.Getenv("SLACK_ICON"); s != "" {
		msg.IconURL = s
	}

	err = slack.Send(webhook, &msg)
	if err != nil {
		log.Fatalf("error sending message: %s", err)
	}

	fmt.Printf("Slack message sent!\n")
}
