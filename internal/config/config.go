package config

import "os"

type Config struct {
	ChannelSecret string
	ChannelToken  string
	UserID        string
	Port          string
}

func New() *Config {
	return &Config{
		ChannelSecret: os.Getenv("LINE_CHANNEL_SECRET"),
		ChannelToken:  os.Getenv("LINE_CHANNEL_TOKEN"),
		UserID:        os.Getenv("LINE_USER_ID"),
		Port:          os.Getenv("PORT"),
	}
}
