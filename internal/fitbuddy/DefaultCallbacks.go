package fitbuddy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartCallback(botInstance *BotInstance, update tgbotapi.Update) error {
	botInstance.SendText(update, "Hello Buddy!")
	return nil
}

func PrintMenuCallback(botInstance *BotInstance, update tgbotapi.Update) error {
	var inline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Мои тренировки", "my exc...")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Мои упражнения", "my trns...")),
	)

	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "Menu: ")
	msgConf.ReplyMarkup = inline

	botInstance.Send(msgConf)

	return nil
}
