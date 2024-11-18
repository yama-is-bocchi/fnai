package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
)

var (
	mainCommand = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "sum",
			Description: "チャンネル履歴を読み取って要約します",
		},
	}
)

func createCommandHandler(restClient restClient) *handler.Mux {
	command := handler.New()
	command.Use(middleware.Logger)
	command.Group(func(router handler.Router) {
		router.Use(middleware.Print("group1"))
		router.Command("/sum", submitLLMHandler(restClient, "OK")) // ここにコマンド入力.
	})
	return command
}

func submitLLMHandler(restClient restClient, content string) handler.CommandHandler {
	return func(event *handler.CommandEvent) error {
		allMessage, err := restClient.GetChannelMessage(event.Channel().ID(), 0)
		if err != nil {
			return fmt.Errorf("failed to get channel message:%wc", err)
		}
		if err := restClient.SendMessage(event.Channel().ID(), "LLMで解析中..."); err != nil {
			return fmt.Errorf("failed to create messages:%w", err)
		}
		allConversion := strings.Join(allMessage, "\n")
		log.Println(allConversion)
		// ここでLLMに送信.
		return event.CreateMessage(discord.MessageCreate{Content: content})
	}
}
