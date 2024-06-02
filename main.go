package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		slog.Error("TELEGRAM_BOT_TOKEN is required")
	}

	bot := NewBot(token)
	var offset int

	for {
		updates, err := bot.getUpdates(offset)
		if err != nil {
			slog.Error("failed to get chat updates:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, update := range updates {
			offset = update.UpdateID + 1

			if !bot.shouldSetFilter(update.Message.Text) {
				continue
			}

			if err := bot.setFilter(update.Message); err != nil {
				slog.Error("failed to set filter:", err)
			}

			// TODO: filter messages by media type.
		}
	}
}
