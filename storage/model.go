package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	IconUri string             `bson:"icon-uri" json:"icon-uri"`
	Name    string             `bson:"name" json:"name"`
	Price   int                `bson:"price" json:"price"`
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	StoreID primitive.ObjectID `bson:"store-id" json:"store-id"`
}

type OrderItemData struct {
	ItemData *Item `bson:"item-data" json:"item-data"`
	Count    int   `bson:"count" json:"count"`
}

type Order struct {
	ItemList  []*OrderItemData    `bson:"items" json:"items"`
	Location  string              `bson:"location" json:"location"`
	Subtotal  int                 `bson:"subtotal" json:"subtotal"`
	CreatedAt primitive.Timestamp `bson:"created-at" json:"created-at"`
	UserData  *User               `bson:"user-data" json:"user-data"`
	StoreData *Store              `bson:"store-data" json:"store-data"`
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
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
