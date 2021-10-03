package server

import (
	"grocery-backend/storage"

	"github.com/gofiber/fiber"
)

type server struct {
	storage storage.Storage
}

func NewGroceryServer(config *ServerConfig, st storage.Storage) Server {
	f := fiber.New()

	f.Get("/", IndexHandler)

	f.Get("/item", GetItemListHandler(st))

	return &server{storage: st}
}

func (s *server) start() error {

	return nil
}
