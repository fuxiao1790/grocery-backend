package main

import (
	"grocery-backend/server"
	"grocery-backend/storage"
	"log"

	"github.com/sirupsen/logrus"
)

const MONGO_DB_URI = "mongodb://localhost:27272"
const DB_NAME = "GROCERY_DB"

func main() {
	storage, err := storage.NewMongoDBStorage(&storage.Config{
		Uri:  MONGO_DB_URI,
		Name: DB_NAME,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	config := &server.Config{
		Port:     443,
		CertFile: "./tls/cert.pem",
		KeyFile:  "./tls/key.pem",
	}
	server := server.NewGroceryServer(config, storage)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
