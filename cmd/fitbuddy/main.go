package main

import (
	"fittgbot/internal/configuration"
	"fittgbot/internal/fitbuddy"
)

func main() {
	conf := configuration.ParseConfiguration("./conf/conf.toml")
	botInstance := fitbuddy.NewBot(conf, fitbuddy.NewDefaultUpdatesConfig())
	cmdMux := fitbuddy.InitCommandMux(botInstance)
	botInstance.StartListening(cmdMux)
}
