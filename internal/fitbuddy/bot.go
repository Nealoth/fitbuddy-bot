package fitbuddy

import (
	"fittgbot/internal/configuration"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type BotInstance struct {
	BotAPI        *tgbotapi.BotAPI
	Configuration configuration.Configuration
	UpdateConfig  tgbotapi.UpdateConfig
}

func NewBot(conf configuration.Configuration, updateConf tgbotapi.UpdateConfig) *BotInstance {
	bot, err := tgbotapi.NewBotAPI(conf.BotInternalConfiguration.BotAuthToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = conf.BotInternalConfiguration.Debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &BotInstance{
		BotAPI:        bot,
		Configuration: conf,
		UpdateConfig:  updateConf,
	}
}

func NewDefaultUpdatesConfig() tgbotapi.UpdateConfig {
	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60
	return updatesConfig
}

func (bot *BotInstance) StartListening(cmdMux *CmdMux) {
	updatesChan, err := bot.BotAPI.GetUpdatesChan(bot.UpdateConfig)

	if err != nil {
		log.Panic("Bot cannot start listen for updates")
	}

	listenForUpdatesFromChan(cmdMux, bot, updatesChan)
}

func listenForUpdatesFromChan(cmdMux *CmdMux, botInstance *BotInstance, updatesChan tgbotapi.UpdatesChannel) {
	for update := range updatesChan {
		if update.Message == nil {
			continue
		}

		go cmdMux.HandleUpdate(botInstance, update)
	}
}
