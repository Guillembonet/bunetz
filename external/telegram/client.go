package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	cfg    Config
	client *tgbotapi.BotAPI
}

type Config struct {
	Token  string
	ChatID int64
}

func NewClient(cfg Config) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create telegram client: %w", err)
	}

	return &Client{
		cfg:    cfg,
		client: client,
	}, nil
}

func (c *Client) SendMessage(name, contact, text string) error {
	msg := tgbotapi.NewMessage(c.cfg.ChatID, fmt.Sprintf("Message from: %s\nContact: %s\n\nMessage: %s", name, contact, text))
	_, err := c.client.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
