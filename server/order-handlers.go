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
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.GetOrderListRes{Orders: nil, Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		orders, err := s.GetOrderList(reqBody.Skip, reqBody.Count, reqBody.UserID)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(&dto.GetOrderListRes{Orders: nil, Error: dto.Err{Error: err}})
			return nil
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(&dto.GetOrderListRes{Orders: orders, Error: dto.Err{Error: nil}})
		return nil
	}
}

func CreateOrderHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.CreateOrderReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			ctx.JSON(&dto.CreateOrderRes{Error: dto.Err{Error: err}})
			return nil
		}

		err = s.CreateOrder(reqBody.Items, reqBody.Location, reqBody.StoreID, reqBody.UserID)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(&dto.CreateOrderRes{Error: dto.Err{Error: err}})
			return nil
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(&dto.CreateOrderRes{Error: dto.Err{Error: nil}})
		return nil
	}
}
