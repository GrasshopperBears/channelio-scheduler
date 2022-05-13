package services

import (
	"fmt"
	"net/http"
	"server/structs"
	"server/texts"
	"server/utils"
)

func PostChannelMessage(blocks []structs.Block, options []string, chatType string, chatId string) error {
	body := utils.HttpBodyBuilder(
		structs.Message{
			Blocks: blocks,
			Options: []string{"silent"},
		},
	)

	client := http.Client{}
	url := fmt.Sprintf("https://api.channel.io/open/v5/%ss/%s/messages", chatType, chatId)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Add("botName", texts.BOT_NAME)
	req.URL.RawQuery = query.Encode()

	utils.SetChannelApiHeader(req)

	_, err = client.Do(req)
	return err
}
