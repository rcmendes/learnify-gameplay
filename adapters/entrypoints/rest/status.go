package rest

import "github.com/gofiber/fiber/v2"

//Status is the API's health check endpoint.
func Status(c *fiber.Ctx) error {
	return c.SendString("Ok")
}
