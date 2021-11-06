package server

import (
	"grocery-backend/storage"

	fiber "github.com/gofiber/fiber/v2"
)

type server struct {
	storage storage.Storage
	app     *fiber.App
}

func NewGroceryServer(config *ServerConfig, st storage.Storage) Server {
	app := fiber.New()

	app.Get("/", IndexHandler)

	app.Get("/item", GetItemListHandler(st))

	return &server{storage: st}
}

func (s *server) Start() error {
	return s.app.Listen("0.0.0.0:8080")
}
