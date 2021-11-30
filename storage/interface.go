package storage

import (
	"grocery-backend/dto"
)

type Storage interface {
	GetItemList(skip int, count int, storeID string) ([]*dto.Item, error)
	GetOrderList(skip int, count int, userID string) ([]*dto.Order, error)
	GetStoreList(skip int, count int) ([]*dto.Store, error)

	CreateItem(*dto.Item) error
	UpdateItem(*dto.Item) error
	DeleteItem(*dto.Item) error

	CreateOrder(orderItems map[string]int, location string, storeID string, userID string) error
	UpdateOrder(*dto.Order) error
	DeleteOrder(*dto.Order) error

	CreateStore(*dto.Store) error
	UpdateStore(*dto.Store) error
	DeleteStore(*dto.Store) error

	CreateUser(*User) error
	GetUser(*User) (*User, error)
}
