package routes

import (
	"github.com/gofiber/fiber/v2"
	"threebaristas.com/purple/app/controllers"
)

func AdminPanelRoutes(a *fiber.App) {
  route := a.Group("/api/v1/admin")
  route.Get("/price", controllers.GetPrice)
}
