package services

import (
	"server/structs"

	"github.com/gofiber/fiber/v2"
)

func GetSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	return ctx.SendStatus(200)
}
