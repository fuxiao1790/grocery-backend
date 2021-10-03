package server

import "github.com/gofiber/fiber"

type IndexResponse struct {
	Data string
}

func IndexHandler(c *fiber.Ctx) {
	c.JSON(&IndexResponse{Data: "hello"})
	c.Status(200)
}
