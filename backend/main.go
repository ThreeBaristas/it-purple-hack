package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/repository"
	"threebaristas.com/purple/pkg/middleware"
	"threebaristas.com/purple/pkg/routes"
)

func main() {
	logger, _ := zap.NewProduction()

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}

	cR := repository.NewCategoriesRepositoryImpl()
	lR := repository.NewLocationsRepositoryImpl()
	pR := repository.NewPostgresPriceRepository(db)

	app := fiber.New()

	middleware.SwaggerMiddleware(app)

	routes.AdminPanelRoutes(app, &cR, &lR, &pR)
	routes.CategoriesRoutes(app, &cR)
	routes.LocationsRoutes(app, &lR)

	logger.Info("Starting web server")
	app.Listen(":3000")
}
