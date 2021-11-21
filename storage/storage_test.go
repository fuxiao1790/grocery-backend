package storage

import (
	"grocery-backend/dto"
	"testing"
)

const MONGO_DB_URI = "mongodb://localhost:27017"
const DB_NAME = "GROCERY_DB"

var TEST_ITEM = &dto.Item{
	IconUri: "test uri",
	Name:    "test name",
	Price:   "test price",
}

var TEST_ORDER = &dto.Order{
	Items:    []dto.Item{*TEST_ITEM},
	Location: "test location",
}

var TEST_STORE = &dto.Store{
	Location: "test location",
	Name:     "test store name",
}

func TestMain(m *testing.M) {
	m.Run()
}

func Test_GetList(t *testing.T) {
	
}