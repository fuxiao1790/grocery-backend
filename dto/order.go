package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Items    []Item             `bson:"items" json:"items"`
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	StoreID  primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type GetOrderListReq struct {
	Skip  int
	Count int
}

type GetOrderListRes struct {
	Orders *[]Order
	Error  error
}
