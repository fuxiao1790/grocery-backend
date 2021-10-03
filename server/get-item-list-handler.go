package server

import (
	"grocery-backend/dto"
	"grocery-backend/storage"
	"net/http"

	"github.com/gofiber/fiber"
)

type GetItemListReq struct {
	Skip  int
	Count int
}

type GetItemListRes struct {
	Items []dto.Item
}

func GetItemListHandler(s storage.Storage) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		reqBody := &GetItemListReq{}
		err := c.BodyParser(reqBody)
		if err != nil {
			c.SendStatus(http.StatusBadRequest)
			return
		}

		s.GetItemList(reqBody.Skip, reqBody.Count)
	}
}
