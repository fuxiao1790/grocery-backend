package server

import (
	"grocery-backend/dto"
	"grocery-backend/storage"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func GetItemListHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.GetItemListReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return err
		}

		items, err := s.GetItemList(reqBody.Skip, reqBody.Count)
		if err != nil {
			ctx.JSON(&dto.GetItemListRes{Items: nil, Error: err})
			return nil
		}

		ctx.JSON(&dto.GetItemListRes{Items: items, Error: nil})
		return nil
	}
}

func NewItemHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.NewItemReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.SendStatus(http.StatusBadRequest)
			return err
		}

		err = s.CreateItem(&dto.Item{
			IconUri: reqBody.IconUri,
			Name:    reqBody.Name,
			Price:   reqBody.Price,
			StoreID: [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		})

		if err != nil {
			ctx.JSON(&dto.NewItemRes{Error: err})
			return nil
		}

		ctx.JSON(&dto.NewItemRes{Error: nil})
		return nil
	}
}
