package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	IconUri string             `bson:"icon-uri" json:"icon-uri"`
	Name    string             `bson:"name" json:"name"`
	Price   string             `bson:"price" json:"price"`
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	StoreID primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type Order struct {
	Items    map[*Item]int      `bson:"items" json:"items"`
	Location string             `bson:"location" json:"location"`
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	UserID   primitive.ObjectID `bson:"user-id" json:"user-id"`
	StoreID  primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type Store struct {
	Location string             `bson:"location" json:"location"`
	Name     string             `bson:"name" json:"name"`
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
}

type User struct {
	Username       string             `bson:"username" json:"username"`
	HashedPassword string             `bson:"hashed-password" json:"hashed-password"`
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
}
