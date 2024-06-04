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
				continue
			}

			if !bot.shoudlDeleteMessage(update.Message) {
				continue
			}

			if err := bot.deleteMessage(update.Message.Chat.ID, update.Message.MessageID); err != nil {
				slog.Error("failed to delete message:", err)
				continue
			}

			if err := bot.sendMessage("Filtered message detected and removed", update.Message.Chat.ID, update.Message.MessageID); err != nil {
				slog.Error("failed to send reply message that was detected and removed:", err)
			}
		}
	}
}
