package fitbuddy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type CmdMux struct {
	botInstance *BotInstance
	storage     ICmdStorage
}

func InitCommandMux(instance *BotInstance) *CmdMux {
	return &CmdMux{
		botInstance: instance,
		storage:     InitCommands(),
	}
}

func (mux *CmdMux) HandleUpdate(botInstance *BotInstance, update tgbotapi.Update) {

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	cb, found := mux.storage.getCmdCallback(update.Message.Text)

	if !found {
		botInstance.TextReply(update, "Unknown command")
		return
	}

	if err := cb(botInstance, update); err != nil {
		log.Printf("[ERR] %s \n", err)
	}

}
