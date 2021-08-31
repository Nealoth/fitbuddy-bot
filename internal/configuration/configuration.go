package configuration

type Configuration struct {
	BotInternalConfiguration botConfiguration `toml:"bot"`
}

type botConfiguration struct {
	BotAuthToken string `env:"BOT_AUTH_TOKEN"`
	Debug        bool   `toml:"debug"`
}
