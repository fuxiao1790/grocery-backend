package storage

import "grocery-backend/dto"

type Storage interface {
	CreateItem(item dto.Item) error
	UpdateItem(item dto.Item) error
	DeleteItem(item dto.Item) error

	CreateOrder(order dto.Order) error
	UpdateOrder(order dto.Order) error
	DeleteOrder(order dto.Order) error

	GetItemList(skip int, count int) (*[]dto.Item, error)
}
