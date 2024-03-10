package middleware

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

  _ "threebaristas.com/purple/docs"
)

func SwaggerMiddleware(a *fiber.App) {
  config := swagger.Config {
    BasePath: "/",
    FilePath: "./docs/swagger.json",
  }
  a.Use(swagger.New(config))
}
