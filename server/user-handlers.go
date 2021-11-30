package server

import (
	"errors"
	"grocery-backend/dto"
	"grocery-backend/storage"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const COST = bcrypt.DefaultCost

var USER_ALREADY_EXIST = errors.New("user already exist")
var USER_DOES_NOT_EXIST = errors.New("user does not exist")

func LoginHandler(st storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// unmarshal request body into struct
		reqBody := &dto.LoginReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.LoginRes{Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		// fetch user from db by user name
		user, err := st.GetUser(&storage.User{Username: reqBody.Username})
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.LoginRes{Error: dto.Err{Error: err}})
			return nil
		}
		// the given username does not exist
		if user == nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.LoginRes{Error: dto.Err{Error: USER_DOES_NOT_EXIST}})
			return nil
		}

		// compare hashed passwrods
		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(reqBody.Password))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.LoginRes{Error: dto.Err{Error: err}})
			return nil
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(dto.LoginRes{
			UserID:   user.ID.Hex(),
			Username: user.Username,
			Error:    dto.Err{Error: nil},
		})

		return nil
	}
}

func RegisterHandler(st storage.Storage) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// unmarshal request body into struct
		reqBody := &dto.RegisterReq{}
		err := ctx.BodyParser(reqBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: CANNOT_PARSE_BODY}})
			return nil
		}

		// fetch user from db by username
		// checking to see if the given username has already been registered
		user, err := st.GetUser(&storage.User{Username: reqBody.Username})
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: nil}})
			return nil
		}
		if user != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: USER_ALREADY_EXIST}})
			return nil
		}

		// generate salted hashed password for security reasons
		saltedHashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), COST)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: err}})
			return nil
		}

		// add the new user to db
		err = st.CreateUser(&storage.User{
			Username:       reqBody.Username,
			HashedPassword: string(saltedHashedPassword),
		})
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: err}})
			return nil
		}

		ctx.Status(http.StatusOK)
		ctx.JSON(dto.RegisterRes{Error: dto.Err{Error: nil}})

		return nil
	}
}
