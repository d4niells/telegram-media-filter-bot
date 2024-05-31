package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const telegramAPI = "https://api.telegram.org/bot"

type UpdateMessageChat struct {
	ID int `json:"id"`
}

type UpdateMessage struct {
	MessageID int               `json:"message_id"`
	Text      string            `json:"text"`
	Chat      UpdateMessageChat `json:"chat"`
	Photo     []any             `json:"photo"`
	Video     any               `json:"video"`
	Document  any               `json:"document"`
}

type Update struct {
	UpdateID int           `json:"update_id"`
	Message  UpdateMessage `json:"message"`
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
