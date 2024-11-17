package main

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/yama-is-bocchi/fnai/bot"
)

type config struct {
	DiscordToken string `env:"DISCORD_TOKEN,required"`
	GuildID      string `env:"GUILD_ID,required"`
	Model        string `env:"MODEL,required"`
}

func main() {
	// envパース
	var config config
	if err := env.Parse(&config); err != nil {
		log.Fatal("failed to parse env file:", err)
	}
	// ボット登録
	originBot, err := bot.New(config.DiscordToken, config.GuildID)
	if err != nil {
		log.Fatal("failed to initialize discord bot:", err)
	}
	// コマンドusage
    
	// llm登録
	// リッスン
}
