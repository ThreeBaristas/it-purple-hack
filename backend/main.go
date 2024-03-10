package main

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"threebaristas.com/purple/pkg/routes"
)

func main()  {
  logger, _ := zap.NewProduction();
  app := fiber.New()

  routes.AdminPanelRoutes(app)

  logger.Info("Starting web server")
  app.Listen(":3000")
}
