package server

import (
	"fmt"
	"grocery-backend/storage"

	fiber "github.com/gofiber/fiber/v2"
)

type server struct {
	storage storage.Storage
	app     *fiber.App
	config  *Config
}

type Config struct {
	PortHttp  int
	PortHttps int
	CertFile  string
	KeyFile   string
}

func NewGroceryServer(config *Config, st storage.Storage) Server {
	app := fiber.New()

	app.Get("/", IndexHandler)

	app.Post("/item/list", GetItemListHandler(st))
	// app.Post("/item/new", GetItemListHandler(st))
	// app.Delete("/item/del", GetItemListHandler(st))

	app.Post("/store/list", GetStoreListHandler(st))
	// app.Post("/store/new", GetItemListHandler(st))
	// app.Delete("/store/del", GetItemListHandler(st))

	app.Post("/order/list", GetOrderListHandler(st))
	// app.Post("/order/new", GetItemListHandler(st))
	// app.Delete("/order/del", GetItemListHandler(st))

	app.Post("/user/login", LoginHandler(st))
	app.Post("/user/register", RegisterHandler(st))

	app.Post("/order/create", CreateOrderHandler(st))
	app.Post("/order/list", GetOrderListHandler(st))

	return &server{storage: st, app: app, config: config}
}

func (s *server) Start() error {
	ch := make(chan error)
	go func() {
		ch <- s.app.Listen(fmt.Sprintf("0.0.0.0:%d", s.config.PortHttp))
	}()

	go func() {
		ch <- s.app.ListenTLS(
			fmt.Sprintf("0.0.0.0:%d", s.config.PortHttps),
			s.config.CertFile,
			s.config.KeyFile,
		)
	}()

	return <-ch
}
