package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
	"threebaristas.com/purple/app/core/services"
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
	err = db.Ping()
	if err != nil {
		logger.Fatal("Could not initialize db connection", zap.Error(err))
	}
	db.SetMaxOpenConns(64)

	cR := repository.NewCategoriesRepositoryImpl()
	lR := repository.NewLocationsRepositoryImpl()
	pR := repository.NewPostgresPriceRepository(db)

	storage := repository.DefaultInlineMappingStorage(db)

	if os.Getenv("GENERATE_STORAGE") == "TRUE" {
		logger.Warn("Started generating storage")
		var discounts []repository.DiscountMappingDTO
		for i := 1; i <= 200; i++ {
			discounts = append(discounts, repository.DiscountMappingDTO{
				SegmentId: int64(i),
				MatrixId:  int64(i),
			})
		}
		storage.SetUpStorage(&repository.SetUpStorageRequest{
			BaselineMatrix: 0,
			Discounts:      discounts,
		})
		logger.Warn("Finished generating storage")
	}

	service := services.NewPriceService(cR, lR, &pR, &storage)
	if os.Getenv("GENERATE_RULES") == "TRUE" {
		logger.Warn("Started generating price rules")
		service.GenerateRules()
		logger.Warn("Finished generating price rules")
	}

	app := fiber.New()

	if os.Getenv("MODE") == "dev" {
		middleware.SwaggerMiddleware(app)
	}

	routes.AdminPanelRoutes(app, cR, lR, &pR, &storage)
	routes.CategoriesRoutes(app, cR)
	routes.LocationsRoutes(app, lR)
	routes.PriceRoutes(app, cR, lR, &pR, &storage)

	p := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	app.Get("/metrics", func(c *fiber.Ctx) error {
		p(c.Context())
		return nil
	})

	logger.Info("Starting web server")
	app.Listen(":3000")
}
