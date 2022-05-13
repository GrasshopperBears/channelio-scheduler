package services

import (
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/structs"
	"server/texts"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func scheduleToString(schedule models.Schedule, idx int) string {
	scheduleString := fmt.Sprintf("[%d] ", idx)

	scheduleDatetime := schedule.Datetime
	scheduleString += fmt.Sprintf("%d/%d/%d ", scheduleDatetime.Year(), scheduleDatetime.Month(), scheduleDatetime.Day())

	if schedule.IsTimeSet {
		scheduleString += fmt.Sprintf("%d:%d ", scheduleDatetime.Hour(), scheduleDatetime.Minute())
	}
	
	scheduleString += schedule.Title
	return scheduleString
}

func GetSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	db := models.DB
	var schedules []models.Schedule
	var blocks []structs.Block

	if result := db.Where("channel_id = ?", event.Entity.ChatId).Order("datetime").Find(&schedules); result.Error != nil {
		log.Println("Error:", result.Error)
		return ctx.SendStatus(500)
	}
	for i := 1; i <= len(schedules); i++ {
		blocks = append(blocks, structs.Block{Type: "text", Value: scheduleToString(schedules[i-1], i)})
	}
	if len(blocks) == 0 {
		blocks = append(blocks, structs.Block{Type: "text", Value: texts.MESSAGE_NO_SCHEDULE})
	}
	
	body := utils.HttpBodyBuilder(
		structs.Message{
			Blocks: blocks,
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
	query.Add("botName", texts.BOT_NAME)
	req.URL.RawQuery = query.Encode()

	utils.SetChannelApiHeader(req)

	_, err = client.Do(req)
	if err != nil {
		log.Println("Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
