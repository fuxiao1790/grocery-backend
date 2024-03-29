package server

import (
	"grocery-backend/dto"
	"grocery-backend/storage"
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

func GetStoreListHandler(s storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		reqBody := &dto.GetStoreListReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.GetItemListRes{Items: nil, Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		items, err := s.GetStoreList(reqBody.Skip, reqBody.Count)
		if err != nil {
			ctx.JSON(&dto.GetStoreListRes{Stores: nil, Error: dto.Err{Error: err}})
			return nil
		}

		ctx.JSON(&dto.GetStoreListRes{Stores: items, Error: dto.Err{Error: err}})
		return nil
	}
}
