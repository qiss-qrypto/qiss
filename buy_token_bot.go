package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func getENVValue(key, defaultValue string) (v string) {
	var found bool
	v, found = os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return v
}

func main() {
	buyBotID := getENVValue("QISS_BOT_ID", "YOUR-BOT-TOKEN")
	bot, err := tgbotapi.NewBotAPI(buyBotID)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("sure:", update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
