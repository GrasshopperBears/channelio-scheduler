package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func initialize() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Environment configuration failed.")
  }
}

func main() {
  initialize()
  app := fiber.New()

  app.Post("/", func(ctx *fiber.Ctx) error {
    if ctx.Query("token") != os.Getenv("TOKEN") {
      log.Println("Request with different token")
      return ctx.SendStatus(400)
    }
    return ctx.SendString("Hello, World!")
  })

  app.Listen(":4000")
}
