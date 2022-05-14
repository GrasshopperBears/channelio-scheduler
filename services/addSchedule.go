package services

import (
	"fmt"
	"log"
	"regexp"
	"server/models"
	"server/structs"
	"server/texts"
	"server/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var dateRegex = "(?:(?P<year>[0-9]{4})/)?(?:(?P<month>[0-9]{1,2})/)(?P<date>[0-9]{1,2}) +"
var timeRegexWord = "(?P<hour>[0-9]{1,2}):(?:(?P<minute>[0-9]{1,2}) +)"

var addRegex = regexp.MustCompile(
	fmt.Sprintf("%s +%s +", texts.SCHEDULER_PREFIX, texts.SCHEDULER_ADD) +
	dateRegex +
	fmt.Sprintf("(?:(%s))?", timeRegexWord) +
	"(?P<title>.+)",
)

func isValidDatetimeInput(year int, month int, date int, hour int, minute int) bool {
	if date == 0 || (month == 0 && year != 0) { return false }
	if (year > 0 && year < 2000) || year > 9999 || month < 0 || month > 12 || date < 0 || date > 31 ||
		hour < 0 || hour > 24 || minute < 0 || minute >= 60 {
			return false
	}
	return true
}

func AddSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	db := models.DB
	match := addRegex.FindStringSubmatch(event.Entity.PlainText)
	parseMap := utils.ParseRegexFind(addRegex, match)

	year, _ := strconv.Atoi(parseMap["year"])
	month, _ := strconv.Atoi(parseMap["month"])
	date, _ := strconv.Atoi(parseMap["date"])
	hour, _ := strconv.Atoi(parseMap["hour"])
	minute, _ := strconv.Atoi(parseMap["minute"])
	title := parseMap["title"]

	if !isValidDatetimeInput(year, month, date, hour, minute) {
		block := structs.Block{Type: "text", Value: texts.MESSAGE_WRONG_FORMAT}
		PostChannelMessage([]structs.Block{block}, []string{"silent"}, event.Entity.ChatType, event.Entity.ChatId)
		return ctx.SendStatus(200)
	}

	today := time.Now()
	if year == 0 {
		year = today.Year()
		if month < int(today.Month()) || month == int(today.Month()) { year++ }
	}

	datetime := time.Date(year, time.Month(month), date, hour, minute, 0, 0, time.Local)

	newSchedule := models.Schedule{}
	newSchedule.ChannelId = event.Entity.ChatId
	newSchedule.Datetime = datetime
	newSchedule.Title = title
	newSchedule.IsTimeSet = parseMap["hour"] != ""

	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newSchedule).Error; err != nil {
			log.Println("Database error: ", err)
			return err
		}
		if err := tx.Where("channel_id = ?", event.Entity.ChatId).Delete(&models.GetScheduleHistory{}).Error; err != nil {
			return err
		}
		return nil
	})

	log.Println("Schedule created")

	block := structs.Block{Type: "text", Value: texts.MESSAGE_ADD_SUCCESS}
	_ = PostChannelMessage([]structs.Block{block}, []string{"silent"}, event.Entity.ChatType, event.Entity.ChatId)

	return ctx.SendStatus(200)
}
