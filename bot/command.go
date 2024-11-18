package bot

import (
	"log"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgo/handler/middleware"
	"github.com/yama-is-bocchi/fnai/llm"
)

var (
	mainCommand = []discord.ApplicationCommandCreate{
		discord.SlashCommandCreate{
			Name:        "sum",
			Description: "チャンネル履歴を読み取って要約します",
		},
	}
)

func createCommandHandler(restClient restClient, llm *llm.LLM) *handler.Mux {
	command := handler.New()
	command.Use(middleware.Logger)
	command.Group(func(router handler.Router) {
		router.Use(middleware.Print("group1"))
		router.Command("/sum", submitLLMHandler(restClient, llm)) // ここにコマンド入力.
	})
	return command
}

func submitLLMHandler(restClient restClient, llm *llm.LLM) handler.CommandHandler {
	return func(event *handler.CommandEvent) error {
		go func() {
			allMessage, err := restClient.GetChannelMessage(event.Channel().ID(), 0,)
			if err != nil {
				log.Printf("failed to get channel message: %v", err)
				restClient.SendMessage(event.Channel().ID(), "メッセージの取得に失敗しました。")
				return
			}

			message, err := llm.SendMessage(strings.Join(allMessage, "\n"))
			if err != nil {
				log.Printf("failed to send message: %v", err)
				restClient.SendMessage(event.Channel().ID(), "解析中にエラーが発生しました。")
				return
			}
			if err := restClient.SendMessage(event.Channel().ID(), message); err != nil {
				log.Printf("failed to send result message: %v", err)
			}
		}()
		return event.CreateMessage(discord.MessageCreate{Content: "解析を開始しました..."})
	}
}
