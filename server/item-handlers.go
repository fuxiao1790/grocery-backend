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
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.GetItemListRes{Items: nil, Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		items, err := s.GetItemList(reqBody.Skip, reqBody.Count, reqBody.StoreID)
		if err != nil {
			ctx.JSON(&dto.GetItemListRes{Items: nil, Error: dto.Err{Error: err}})
			return nil
		}

		ctx.JSON(&dto.GetItemListRes{Items: items, Error: dto.Err{Error: nil}})
		return nil
	}
}

func NewItemHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.NewItemReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.GetItemListRes{Items: nil, Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		err = s.CreateItem(&dto.Item{
			IconUri: reqBody.IconUri,
			Name:    reqBody.Name,
			Price:   reqBody.Price,
		})

		if err != nil {
			ctx.JSON(&dto.NewItemRes{Error: dto.Err{Error: err}})
			return nil
		}

		ctx.JSON(&dto.NewItemRes{Error: dto.Err{Error: err}})
		return nil
	}
}
