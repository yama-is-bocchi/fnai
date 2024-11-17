package bot

import (
	"fmt"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
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
	client, err := disgo.New(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord client:%w", err)
	}
	return &originBot{
		client:  client,
		guildID: parsedGuildID,
	}, nil
}
