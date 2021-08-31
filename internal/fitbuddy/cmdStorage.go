package fitbuddy

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type callbackFunc func(botInstance *BotInstance, update tgbotapi.Update) error

type ICmdStorage interface {
	addCallback(cmd string, callback callbackFunc)
	getCmdCallback(cmd string) (callbackFunc, bool)
}

type CmdStorage struct {
	storage map[string]callbackFunc
}

func InitCmdStorage() *CmdStorage {
	return &CmdStorage{
		storage: make(map[string]callbackFunc),
	}
}

func (cs *CmdStorage) addCallback(cmd string, callback callbackFunc) {
	cs.storage[cmd] = callback
}

func (cs *CmdStorage) getCmdCallback(cmd string) (callbackFunc, bool) {
	callback, found := cs.storage[cmd]
	return callback, found
}
