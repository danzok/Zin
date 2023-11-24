package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6751654636:AAFJrNrTEBEfR3JMfrKzp2GLzxKYjTYYlW0")
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

			if !update.Message.IsCommand() { // ignore any non-command Messages
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			// Extract the command from the Message.
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "start":
				msg.Text = "Olá  , sou um bot que converte cookie plain ou base64  para o formato json , com a ajuda do site accovod.com, apenas me envie o comando  /format aqui é seu cookie"
			case "status":
				msg.Text = "I'm ok. and running hehe"
			case "format":
				msg.Text = "comando format criado"
			default:
				msg.Text = "I don't know that command"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
