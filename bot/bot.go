package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
)

type originBot struct {
	guildID snowflake.ID
	client  bot.Client
}

func New(token, guildID string) (*originBot, error) {
	parsedGuildID, err := snowflake.Parse(guildID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse to guildID:%w", err)
	}
	restClient := NewRestClient(token)
	client, err := disgo.New(token,
		bot.WithDefaultGateway(),
		bot.WithEventListeners(createCommandHandler(restClient)))
	if err != nil {
		return nil, fmt.Errorf("failed to create discord client:%w", err)
	}
	return &originBot{
		client:  client,
		guildID: parsedGuildID,
	}, nil
}

func (ogBot originBot) Listen() error {
	if err := handler.SyncCommands(ogBot.client, mainCommand, []snowflake.ID{ogBot.guildID}); err != nil {
		return fmt.Errorf("failed to set sync command:%w", err)
	}
	defer ogBot.client.Close(context.TODO())
	if err := ogBot.client.OpenGateway(context.TODO()); err != nil {
		return fmt.Errorf("failed to open gateway:%w", err)
	}
	log.Println("discord bot running...")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
	return nil
}
