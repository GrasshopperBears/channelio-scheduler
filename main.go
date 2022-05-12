package main

import (
	"fmt"
	"log"
	"os"
	"server/models"
	"server/services"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func initialize() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Environment configuration failed.")
  }

  models.ConnectDatabase()
}

func main() {
  initialize()
  app := fiber.New()
  port := os.Getenv("PORT")
  if len(port) == 0 { port = "4000" }

  app.Post("/", func(ctx *fiber.Ctx) error {
    if ctx.Query("token") != os.Getenv("TOKEN") {
      log.Println("Request with different token")
      return ctx.SendStatus(200)
    }
    return services.HookEntryHandler(ctx)
  })

  if err := app.Listen(fmt.Sprint(":", port)); err != nil {
    log.Fatal("Server start error", err)
  }
}
