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
			slog.Info("Chat message: %+v\n", update)
		}
	}
}
