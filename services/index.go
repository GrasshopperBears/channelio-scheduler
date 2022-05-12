package services

import (
	"log"
	"server/structs"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const MSG_PREFIX = "-일정"

func HookEntryHandler(ctx *fiber.Ctx) error {
	event := new(structs.WebhookEvent)

	if err := ctx.BodyParser(event); err != nil {
		log.Println("Cannot parse body.")
		return ctx.SendStatus(500)
	}

	event.Entity.PlainText = strings.TrimSpace(event.Entity.PlainText)
	if !strings.HasPrefix(event.Entity.PlainText, MSG_PREFIX) {
		return ctx.SendStatus(200)
	}

	return ctx.SendString("Hello, World!")
}
