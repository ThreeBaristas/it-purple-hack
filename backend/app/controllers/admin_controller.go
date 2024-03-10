package controllers

import "github.com/gofiber/fiber/v2"

// GetPrice func gets a price for given category_id & location_id
// @Tags admin
// @Produce json
func GetPrice(c *fiber.Ctx) error {
  return c.SendString("hello")
}
