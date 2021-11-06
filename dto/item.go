package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	IconUri string             `bson:"icon-uri" json:"icon-uri"`
	Name    string             `bson:"name" json:"name"`
	Price   string             `bson:"price" json:"price"`
	ID      primitive.ObjectID `bson:"id" json:"id"`
}

type Order struct {
	Items    []Item             `bson:"items" json:"items"`
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"id" json:"id"`
}
