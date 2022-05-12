package services

import (
	"server/models"
	"server/structs"

	"github.com/gofiber/fiber/v2"
)

func GetSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	db := models.DB
	var schedules []models.Schedule

	result := db.Where("channel_id = ?", event.Entity.ChatId).Find(&schedules)
	println(result.RowsAffected)

	return ctx.SendStatus(200)
}
