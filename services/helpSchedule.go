package services

import (
	"log"
	"server/structs"
	"server/texts"

	"github.com/gofiber/fiber/v2"
)

func HelpSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	block := structs.Block{Type: "text", Value: texts.MESSAGE_HELP}

	if err := PostChannelMessage([]structs.Block{block}, []string{"silent"},
																event.Entity.ChatType, event.Entity.ChatId); err != nil {
		log.Println("API error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
