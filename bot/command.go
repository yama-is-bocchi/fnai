package bot

import (
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

func createCommandHandler() *handler.Mux {
	command := handler.New()
	command.Use(middleware.Logger)
	command.Group(func(router handler.Router) {
		router.Use(middleware.Print("group1"))
		router.Command("/sum", handleContent("OK")) // ここにコマンド入力.
	})
	return command
}

func handleContent(content string) handler.CommandHandler {
	return func(event *handler.CommandEvent) error {
		return event.CreateMessage(discord.MessageCreate{Content: content})
	}
}
