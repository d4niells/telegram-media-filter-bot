package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type Chat struct {
	ID int `json:"id"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	Text      string `json:"text"`
	Chat      Chat   `json:"chat"`
	Photo     []any  `json:"photo"`
	Video     any    `json:"video"`
	Document  any    `json:"document"`
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Bot struct {
	token      string
	filterType string
	mu         sync.RWMutex
}

func NewBot(token string) *Bot {
	return &Bot{token: token}
}

func (b *Bot) withToken(path string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s%s", b.token, path)
}

func (b *Bot) getUpdates(offset int) ([]Update, error) {
	url := b.withToken(fmt.Sprintf("/getUpdates?offset=%d&timeout=60", offset))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result struct {
		OK     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Result, nil
}

func (b *Bot) sendMessage(text string, chatID, replyToMessageID int) error {
	msg := map[string]any{
		"chat_id":             chatID,
		"text":                text,
		"reply_to_message_id": replyToMessageID,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := http.Post(b.withToken("/sendMessage"), "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (b *Bot) shouldSetFilter(msg string) bool {
	return msg != "" && strings.HasPrefix(msg, "/setfilter")
}

func (b *Bot) setFilter(msg Message) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	filterType := strings.TrimSpace(strings.TrimPrefix(msg.Text, "/setfilter"))
	if filterType != "photo" && filterType != "video" && filterType != "document" && filterType != "text" && filterType != "none" {
		return b.sendMessage("Invalid filter type. Use photo, video, document, text or none.", msg.Chat.ID, msg.MessageID)
	}
	b.filterType = filterType
	return b.sendMessage(fmt.Sprintf("Filter type set to %s", filterType), msg.Chat.ID, msg.MessageID)
}

func (b *Bot) shoudlDeleteMessage(msg Message) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	switch b.filterType {
	case "photo":
		return msg.Photo != nil
	case "video":
		return msg.Video != nil
	case "document":
		return msg.Document != nil
	case "text":
		return msg.Text != ""
	default:
		return false
	}
}

func (b *Bot) deleteMessage(chatID, messageID int) error {
	msg := map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := http.Post(b.withToken("/deleteMessage"), "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
