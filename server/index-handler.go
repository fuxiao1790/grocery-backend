package server

import fiber "github.com/gofiber/fiber/v2"

type IndexResponse struct {
	Data string
}

func IndexHandler(c *fiber.Ctx) error {
	c.JSON(&IndexResponse{Data: "hello"})
	c.Status(200)
	return nil
}
