package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Set up a channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the bot in a separate goroutine
	go func() {
		if err := b.Run(); err != nil {
			log.Printf("Error running bot: %v", err)
			stop <- os.Interrupt // Signal the main goroutine to stop
		}
	}()

	// Wait for an interrupt signal
	<-stop

	// Perform cleanup operations here (if any)
	log.Println("Shutting down gracefully...")
	// Add any necessary cleanup code here

	log.Println("Bot has been shut down")
}
