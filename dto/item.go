package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	IconUri string             `bson:"icon-uri" json:"icon-uri"`
	Name    string             `bson:"name" json:"name"`
	Price   string             `bson:"price" json:"price"`
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	StoreID primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type GetItemListReq struct {
	Skip    int
	Count   int
	StoreID string
}

type GetItemListRes struct {
	Items *[]Item
	Error error
}

type NewItemReq struct {
	IconUri string             `bson:"icon-uri" json:"icon-uri"`
	Name    string             `bson:"name" json:"name"`
	Price   string             `bson:"price" json:"price"`
	StoreID primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type NewItemRes struct {
	Error error
}
