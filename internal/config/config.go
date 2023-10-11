package config

import (
	"log"
	"os"
)

type Config struct {
	BotToken string
}

func MustLoad() *Config {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("BOT_TOKEN env is not set")
	}

	return &Config{
		BotToken: token,
	}
}
