package main

import (
	"grocery-backend/server"
	"grocery-backend/storage"
	"log"

	"github.com/sirupsen/logrus"
)

const MONGO_DB_URI = "mongodb://localhost:27017"
const DB_NAME = "GROCERY_DB"

func main() {
	storage, err := storage.NewMongoDBStorage(&storage.Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	server := server.NewGroceryServer(&server.Config{Port: 8080}, storage)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
