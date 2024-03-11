package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/repository"
	"threebaristas.com/purple/pkg/middleware"
	"threebaristas.com/purple/pkg/routes"
)

func main() {
	logger, _ := zap.NewProduction()

  postgres_host := os.Getenv("DB_HOST")
  if postgres_host == "" {
    postgres_host = "localhost"
  }
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@%s:5432/postgres?sslmode=disable", postgres_host))
	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}
  time.Sleep(1 * time.Second)
	err = db.Ping()
	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}

	cR := repository.NewCategoriesRepositoryImpl()
	lR := repository.NewLocationsRepositoryImpl()
	pR := repository.NewPostgresPriceRepository(db)

	app := fiber.New()

  if os.Getenv("MODE") == "dev" {
    middleware.SwaggerMiddleware(app)
  }

	routes.AdminPanelRoutes(app, &cR, &lR, &pR)
	routes.CategoriesRoutes(app, &cR)
	routes.LocationsRoutes(app, &lR)
	routes.PriceRoutes(app, &cR, &lR, &pR)

	logger.Info("Starting web server")
	app.Listen(":3000")
}
