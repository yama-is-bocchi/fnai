package main

import (
	"log"
	"net/url"

	"github.com/caarlos0/env"
	"github.com/yama-is-bocchi/fnai/bot"
	"github.com/yama-is-bocchi/fnai/llm"
)

type config struct {
	DiscordToken  string `env:"DISCORD_TOKEN,required"`
	GuildID       string `env:"GUILD_ID,required"`
	Model         string `env:"MODEL,required"`
	ModelFilePath string `env:"MODEL_FILE_PATH,required"`
	LlmIP         string `env:"LLM_IP,required"`
}

func main() {
	// envパース
	var config config
	if err := env.Parse(&config); err != nil {
		log.Fatal("failed to parse env file:", err)
	}
	// llm登録
	url, err := url.Parse(config.LlmIP)
	if err != nil {
		log.Fatal("failed to parse url:", err)
	}
	llm, err := llm.New(url, config.ModelFilePath)
	if err != nil {
		log.Fatal("failed to register llm:", err)
	}
	if err := llm.CreateModel(config.Model, true); err != nil {
		log.Fatal("failed to create model:", err)
	}
	// ボット登録
	originBot, err := bot.NewBot(config.DiscordToken, config.GuildID, &llm)
	if err != nil {
		log.Fatal("failed to initialize discord bot:", err)
	}
	// コマンドusage
	if err := originBot.Listen(); err != nil {
		log.Fatal("failed to listen command:", err)
	}
}
