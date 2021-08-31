package fitbuddy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (bot *BotInstance) Send(msg tgbotapi.MessageConfig) {
	_, err := bot.BotAPI.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func (bot *BotInstance) SendText(update tgbotapi.Update, msg string) {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	bot.Send(msgConf)
}

func (bot *BotInstance) TextReply(update tgbotapi.Update, msg string) {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
	msgConf.ReplyToMessageID = update.Message.MessageID
	bot.Send(msgConf)
}
