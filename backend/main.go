package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/pkg/middleware"
	"threebaristas.com/purple/pkg/routes"
)

func main() {
	logger, _ := zap.NewProduction()

  _, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable");
  if err != nil {
    logger.Error("Could not initialize db connection", zap.Error(err))
  }

	app := fiber.New()

	middleware.SwaggerMiddleware(app)

	routes.AdminPanelRoutes(app)
	routes.CategoriesRoutes(app)
	routes.LocationsRoutes(app)

	logger.Info("Starting web server")
	app.Listen(":3000")
}
