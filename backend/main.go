package main

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main()  {
  logger, _ := zap.NewProduction();
  app := fiber.New()

  app.Get("/", func (c fiber.Ctx) error {
    defer logger.Sync();
    logger.Info("Hello from / endpoint")
    return c.SendString("Hello there!")
  })

  logger.Info("Starting web server")
  app.Listen(":3000")
}
