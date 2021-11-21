package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
}

type GetStoreListReq struct {
	Skip  int
	Count int
}

type GetStoreListRes struct {
	Stores *[]Store
	Error  error
}
