package storage

import "grocery-backend/dto"

type storage struct{}

func NewMongoDBStorage() Storage {
	return &storage{}
}

func (*storage) CreateItem(item dto.Item) error {
	return nil
}

func (*storage) UpdateItem(item dto.Item) error {
	return nil
}

func (*storage) DeleteItem(item dto.Item) error {
	return nil
}

func (*storage) GetItemList(skip int, count int) ([]dto.Item, error) {
	return []dto.Item{}, nil
}
