package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Name     string             `bson:"name" json:"name"`
}

type GetStoreListReq struct {
	Skip  int `bson:"skip" json:"skip"`
	Count int `bson:"count" json:"count"`
}

type GetStoreListRes struct {
	Stores []*Store `bson:"stores" json:"stores"`
	Error  error    `bson:"error" json:"error"`
}
