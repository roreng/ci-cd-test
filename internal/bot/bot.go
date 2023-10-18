package bot

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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

			if umsg.Text == "/start" {
				msg := tgbotapi.NewMessage(umsg.Chat.ID, "Welcome!")
				msg.ReplyToMessageID = umsg.MessageID

				_, err := b.tg.Send(msg)
				if err != nil {
					log.Printf("can't send message: %s", err)
					continue
				}

				continue
			}

			if umsg.Text == "/ping" {
				msg := tgbotapi.NewMessage(umsg.Chat.ID, "🏓 pong")
				msg.ReplyToMessageID = umsg.MessageID

				_, err := b.tg.Send(msg)
				if err != nil {
					log.Printf("can't send message: %s", err)
					continue
				}

				continue
			}

			if strings.HasPrefix(umsg.Text, "/sum") {
				args := strings.Split(umsg.Text, " ")
				if len(args) != 3 {
					log.Print("bad command format")
					continue
				}

				first, err := strconv.ParseInt(args[1], 10, 64)
				if err != nil {
					log.Printf("bad first argument: %s", err)
					continue
				}

				second, err := strconv.ParseInt(args[2], 10, 64)
				if err != nil {
					log.Printf("bad first argument: %s", err)
					continue
				}

				sum := first + second
				msgText := fmt.Sprintf("(%v + %v) = %v", first, second, sum)

				msg := tgbotapi.NewMessage(umsg.Chat.ID, msgText)
				msg.ReplyToMessageID = umsg.MessageID

				_, err = b.tg.Send(msg)
				if err != nil {
					log.Printf("can't send message: %s", err)
				}

				continue
			}
		}
	}
}
