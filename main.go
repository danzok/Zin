package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	/*
		//https://pkg.go.dev/github.com/Valgard/godotenv#section-readme
		dotenv := godotenv.New()
		if err := dotenv.Load(".env"); err != nil {
			panic(err)
		}

		tokenBOt := os.Getenv("TOKEN")

		fmt.Println(tokenBOt)
	*/
	bot, err := tgbotapi.NewBotAPI("your token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
