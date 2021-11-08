package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"id" json:"id"`
}

type GetStoreListReq struct {
	Skip  int
	Count int
}

type GetStoreListRes struct {
	Stores *[]Store
	Error  error
}
