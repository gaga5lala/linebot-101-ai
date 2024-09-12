package bot

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"linebot-101/internal/config"
)

type Bot struct {
	client *linebot.Client
	config *config.Config
}

func New(cfg *config.Config) (*Bot, error) {
	client, err := linebot.New(cfg.ChannelSecret, cfg.ChannelToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		client: client,
		config: cfg,
	}, nil
}

func (b *Bot) Run() error {
	http.HandleFunc("/callback", b.handleCallback)
	return http.ListenAndServe(":"+b.config.Port, nil)
}

func (b *Bot) handleCallback(w http.ResponseWriter, r *http.Request) {
	events, err := b.client.ParseRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, event := range events {
		if event.Source.UserID != b.config.UserID {
			continue // Only process messages from authorized user
		}

		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := b.handleTextMessage(event.ReplyToken, message); err != nil {
					fmt.Printf("Error handling message: %v\n", err)
				}
			}
		}
	}
}

func (b *Bot) handleTextMessage(replyToken string, message *linebot.TextMessage) error {
	reply := fmt.Sprintf("Received: %s", message.Text)
	_, err := b.client.ReplyMessage(replyToken, linebot.NewTextMessage(reply)).Do()
	return err
}

func (b *Bot) SendInitialMessage() error {
	message := "Hola! This is from a Line Bot. Produced with Cursor AI!"
	_, err := b.client.PushMessage(b.config.UserID, linebot.NewTextMessage(message)).Do()
	return err
}

func (b *Bot) SendReport(report string) error {
	_, err := b.client.PushMessage(b.config.UserID, linebot.NewTextMessage(report)).Do()
	return err
}
