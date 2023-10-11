package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	tg *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	log.Println("Bot initialization...")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{tg: bot}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.tg.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.tg.GetUpdatesChan(u)

	for update := range updates {
		if umsg := update.Message; umsg != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if umsg.Text == "/ping" {
				msg := tgbotapi.NewMessage(umsg.Chat.ID, "pong")
				msg.ReplyToMessageID = umsg.MessageID

				_, err := b.tg.Send(msg)
				if err != nil {
					log.Printf("can't send message: %s", err)
				}
			}
		}
	}
}
