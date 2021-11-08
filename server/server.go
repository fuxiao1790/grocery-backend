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

	app.Get("/item/list", GetItemListHandler(st))
	// app.Post("/item/new", GetItemListHandler(st))
	// app.Delete("/item/del", GetItemListHandler(st))

	app.Get("/store/list", GetStoreListHandler(st))
	// app.Post("/store/new", GetItemListHandler(st))
	// app.Delete("/store/del", GetItemListHandler(st))

	app.Get("/order/list", GetOrderListHandler(st))
	// app.Post("/order/new", GetItemListHandler(st))
	// app.Delete("/order/del", GetItemListHandler(st))

	return &server{storage: st}
}

func (s *server) Start() error {
	return s.app.Listen("0.0.0.0:8080")
}
