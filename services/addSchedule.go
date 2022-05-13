package services

import (
	"log"
	"regexp"
	"server/models"
	"server/structs"
	"server/texts"
	"server/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var regex = regexp.MustCompile(texts.SCHEDULER_PREFIX + " +" + texts.SCHEDULER_ADD + "(?: +(?P<year>[0-9]{4})년)? +(?P<month>[0-9]{1,2})월 +(?P<date>[0-9]{1,2})일 +(?P<hour>[0-9]{1,2})시 +(?P<minute>[0-9]{1,2})분 (?P<title>.+)")

func AddSchedule(ctx *fiber.Ctx, event *structs.WebhookEvent) error {
	match := regex.FindStringSubmatch(event.Entity.PlainText)
	parseMap := utils.ParseRegexFind(regex, match)

	year, _ := strconv.Atoi(parseMap["year"])
	month, _ := strconv.Atoi(parseMap["month"])
	date, _ := strconv.Atoi(parseMap["date"])
	hour, _ := strconv.Atoi(parseMap["hour"])
	minute, _ := strconv.Atoi(parseMap["minute"])
	title := parseMap["title"]

	datetime := time.Date(year, time.Month(month), date, hour, minute, 0, 0, time.Local)

	newSchedule := models.Schedule{}
	newSchedule.ChannelId = event.Entity.ChatId
	newSchedule.Datetime = datetime
	newSchedule.Title = title
	newSchedule.IsTimeSet = true

	result := models.DB.Create(&newSchedule)
	
	if result.Error != nil {
		log.Println("Error: ", result.Error)
		return ctx.SendStatus(500)
	}
	log.Println("Schedule created")

	return ctx.SendStatus(200)
}
