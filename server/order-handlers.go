package server

import (
	"grocery-backend/dto"
	"grocery-backend/storage"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func GetOrderListHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.GetOrderListReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return nil
		}

		items, err := s.GetOrderList(reqBody.Skip, reqBody.Count)
		if err != nil {
			ctx.JSON(&dto.GetOrderListRes{Orders: nil, Error: err})
			return nil
		}

		ctx.JSON(&dto.GetOrderListRes{Orders: items, Error: nil})
		return nil
	}
}
