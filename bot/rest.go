package bot

import (
	"fmt"
	"log"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/snowflake/v2"
)

type restClient struct {
	client       rest.Rest
	excludeNames []string
}

func NewRestClient(token string, excludeName ...string) restClient {
	return restClient{
		client:       rest.New(rest.NewClient(token)),
		excludeNames: excludeName,
	}
}

func (restClient restClient) GetChannelMessage(channelID snowflake.ID, limit int) ([]string, error) {
	messages, err := restClient.client.GetMessages(channelID, channelID, channelID, channelID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages:%w", err)
	}
	var allMessageArray []string
	for _, msg := range messages {
		excludeFlag := false
		for _, name := range restClient.excludeNames {
			if name == msg.Author.Username {
				excludeFlag = true
				break
			}
		}
		if excludeFlag {
			continue
		}
		log.Println(msg.Content)
		allMessageArray = append(allMessageArray, msg.Content)
	}
	return allMessageArray, nil
}

func (restClient restClient) SendMessage(channelID snowflake.ID, message string) error {
	_, err := restClient.client.CreateMessage(channelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	return err
}
