package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/pkg/middleware"
	"threebaristas.com/purple/pkg/routes"
)

func main() {
	logger, _ := zap.NewProduction()
	app := fiber.New()

	middleware.SwaggerMiddleware(app)

	routes.AdminPanelRoutes(app)

	logger.Info("Starting web server")
	app.Listen(":3000")
}
