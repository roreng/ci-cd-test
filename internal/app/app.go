package app

import (
	"log"

	"github.com/roreng/ci-cd-test/internal/bot"
	"github.com/roreng/ci-cd-test/internal/config"
)

func Run() {
	cfg := config.MustLoad()

	bot, err := bot.NewBot(cfg.BotToken)
	if err != nil {
		log.Fatalf("failed bot init: %s", err)
	}

	bot.Start()
}
