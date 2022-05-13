package services

import (
	"fmt"
	"log"
	"server/models"
	"server/structs"
	"server/texts"
	"time"

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
	var ids []string

	if result := db.Where("channel_id = ? AND datetime > ?", event.Entity.ChatId, time.Now()).Order("datetime").Find(&schedules); result.Error != nil {
		log.Println("Error:", result.Error)
		return ctx.SendStatus(500)
	}
	for i := 0; i < len(schedules); i++ {
		blocks = append(blocks, structs.Block{Type: "text", Value: scheduleToString(schedules[i], i+1)})
		ids = append(ids, schedules[i].ID.String())
	}
	if len(blocks) == 0 {
		blocks = append(blocks, structs.Block{Type: "text", Value: texts.MESSAGE_NO_SCHEDULE})
	}

	getScheduleHistory := models.GetScheduleHistory{}
	result := db.Where("person_id = ? AND channel_id = ?", event.Entity.PersonId, event.Entity.ChatId).Limit(1).Find(&getScheduleHistory)
	
	if result.Error != nil {
		log.Println("Database error:", result.Error)
		return ctx.SendStatus(500)
	}
	
	getScheduleHistory.Result = ids

	if result.RowsAffected == 0 {
		getScheduleHistory.PersonId = event.Entity.PersonId
		getScheduleHistory.ChannelId = event.Entity.ChatId
		if result = db.Create(&getScheduleHistory); result.Error != nil {
			log.Println("Database error:", result.Error)
			return ctx.SendStatus(500)
		}
	} else if result = db.Save(&getScheduleHistory); result.Error != nil {
		log.Println("Database error:", result.Error)
		return ctx.SendStatus(500)
	}

	if err := PostChannelMessage(blocks, []string{"silent"},
																event.Entity.ChatType, event.Entity.ChatId); err != nil {
		log.Println("Error:", err)
		return ctx.SendStatus(500)
	}

	return ctx.SendStatus(200)
}
