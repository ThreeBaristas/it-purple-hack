package controllers

import "github.com/gofiber/fiber/v3"

// GetPrice func gets a price for given category_id & location_id
// @Tags admin
// @Produce json
// @Success 200 {number}
func GetPrice(c fiber.Ctx) error {
  return c.SendString("hello")
}
