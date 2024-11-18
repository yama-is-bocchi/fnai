package bot

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/snowflake/v2"
)

type restClient struct {
	client rest.Rest
}

func NewRestClient(token string) restClient {
	return restClient{
		client: rest.New(rest.NewClient(token)),
	}
}

func (restClient restClient) GetChannelMessage(channelID snowflake.ID, limit int) ([]string, error) {
	messages, err := restClient.client.GetMessages(channelID, channelID, channelID, channelID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages:%w", err)
	}
	var allMessageArray []string
	for _, msg := range messages {
		allMessageArray = append(allMessageArray, msg.Content)
	}
	return allMessageArray, nil
}

func (restClient restClient) SendMessage(channelID snowflake.ID, message string) error {
	_, err := restClient.client.CreateMessage(channelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	return err
}
