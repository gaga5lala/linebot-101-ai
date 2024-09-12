package main

import (
	"log"

	"github.com/joho/godotenv"
	"linebot-101/internal/bot"
	"linebot-101/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.New()
	b, err := bot.New(cfg)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	if err := b.SendInitialMessage(); err != nil {
		log.Printf("Error sending initial message: %v", err)
	}

	if err := b.Run(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
