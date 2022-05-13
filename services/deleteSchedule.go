package services

import (
	"errors"
	"log"
	"regexp"
	"server/models"
	"server/structs"
	"server/texts"
	"server/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var deleteRegex = regexp.MustCompile(texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_DELETE + " (?P<idx>[0-9]+)")

func DeleteSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	db := models.DB
	match := deleteRegex.FindStringSubmatch(event.Entity.PlainText)
	parseMap := utils.ParseRegexFind(deleteRegex, match)
	deleteIdx, err := strconv.Atoi(parseMap["idx"])

	if err != nil {
		block := structs.Block{Type: "text", Value: texts.MESSAGE_WRONG_FORMAT}
		if err := PostChannelMessage([]structs.Block{block}, []string{"silent"},
															event.Entity.ChatType, event.Entity.ChatId); err != nil {
			log.Println("API error:", err)
		}
		return ctx.SendStatus(500)
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		var getScheduleHistory models.GetScheduleHistory
		if err := tx.Where("channel_id = ? AND person_id = ?", event.Entity.ChatId, event.Entity.PersonId).Limit(1).Find(&getScheduleHistory).Error; err != nil {
			log.Println("Database error:", err)
			return err
		}

		if len(getScheduleHistory.Result) == 0 {
			block := structs.Block{Type: "text", Value: texts.MESSAGE_DELETE_BEFORE_GET}
			if err := PostChannelMessage([]structs.Block{block}, []string{"silent"},
																event.Entity.ChatType, event.Entity.ChatId); err != nil {
				return errors.New("")
			}
			return errors.New("")
		}
		if len(getScheduleHistory.Result) < deleteIdx {
			block := structs.Block{Type: "text", Value: texts.MESSAGE_SCHEDULE_NOT_FOUND}
			if err := PostChannelMessage([]structs.Block{block}, []string{"silent"},
																event.Entity.ChatType, event.Entity.ChatId); err != nil {
				return errors.New("")
			}
			return errors.New("")
		}
		deleteScheduleId := getScheduleHistory.Result[deleteIdx - 1]
		deleteSchedule := models.Schedule{}

		result := tx.Where("id = ? and channel_id = ?", uuid.MustParse(deleteScheduleId), event.Entity.ChatId).Limit(1).Find(&deleteSchedule)
		if result.Error != nil {
			return result.Error
		} else if result.RowsAffected == 0 {
			block := structs.Block{Type: "text", Value: texts.MESSAGE_DELETE_BEFORE_GET}
			if err := PostChannelMessage([]structs.Block{block}, []string{"silent"},
																event.Entity.ChatType, event.Entity.ChatId); err != nil {
				return errors.New("")
			}
			return result.Error
		}

		if err := tx.Delete(&deleteSchedule).Error; err != nil {
			log.Println("Database error:", err)
			return err
		}
	
		if err := tx.Delete(&getScheduleHistory).Error; err != nil {
			log.Println("Database error:", err)
			return err
		}
		return nil
	})

	if err != nil {
		return ctx.SendStatus(500)
	}

	println("Schedule deleted")

	return ctx.SendStatus(200)
}
