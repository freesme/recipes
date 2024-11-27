package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// Hello handle api status
func Hello(c *fiber.Ctx) error {
	log.Println("Request received at /api")
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
