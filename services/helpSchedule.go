package services

import (
	"fmt"
	"log"
	"net/http"
	"server/structs"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func HelpSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	block := structs.Block{Type: "text", Value: utils.MESSAGE_HELP}
	body := utils.HttpBodyBuilder(
		structs.Message{
			Blocks: []structs.Block{block},
			Options: []string{"silent"},
		},
	)

	client := http.Client{}
	url := fmt.Sprintf("https://api.channel.io/open/v5/%ss/%s/messages", event.Entity.ChatType, event.Entity.ChatId)
	req, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Println("Error:", err)
		return ctx.SendStatus(500)
	}

	query := req.URL.Query()
	query.Add("botName", "도움말봇")
	req.URL.RawQuery = query.Encode()

	utils.SetChannelApiHeader(req)

	_, err = client.Do(req)
	if err != nil {
		log.Println("Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
