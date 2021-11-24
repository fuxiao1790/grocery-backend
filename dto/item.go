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
	Skip    int    `bson:"skip" json:"skip"`
	Count   int    `bson:"count" json:"count"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type GetItemListRes struct {
	Items []*Item `bson:"items" json:"items"`
	Error error   `bson:"error" json:"error"`
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
