package storage

import (
	"context"
	"grocery-backend/dto"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const ITEM_COLLECTION_NAME = "ITEMS"
const ORDER_COLLECTION_NAME = "ORDERS"
const STORES_COLLECTION_NAME = "STORES"

type storage struct {
	client           *mongo.Client
	itemsCollection  *mongo.Collection
	ordersCollection *mongo.Collection
	storesCollection *mongo.Collection
}

type Config struct {
	Uri  string
	Name string
}

func NewMongoDBStorage(config *Config) (Storage, error) {
	re := &storage{}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Uri))
		if err != nil {
			return nil, err
		}

		re.client = client
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := re.client.Ping(ctx, readpref.Primary())
		if err != nil {
			return nil, err
		}
	}

	re.itemsCollection = re.client.Database(config.Name).Collection(ITEM_COLLECTION_NAME)
	re.ordersCollection = re.client.Database(config.Name).Collection(ORDER_COLLECTION_NAME)
	re.storesCollection = re.client.Database(config.Name).Collection(STORES_COLLECTION_NAME)

	return re, nil
}

func (s *storage) CreateItem(item *dto.Item) error {
	logrus.Info("Create Item")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.itemsCollection.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	logrus.Info(res)

	return nil
}

func (s *storage) UpdateItem(item *dto.Item) error {
	return nil
}

func (s *storage) DeleteItem(item *dto.Item) error {
	logrus.Info("Delete Item")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := s.itemsCollection.FindOneAndDelete(ctx, bson.M{"_id": item.ID})

	logrus.Info(res)
	return nil
}

func (s *storage) GetItemList(skip int, count int, storeID primitive.ObjectID) (*[]dto.Item, error) {
	logrus.Info("Get Item List")

	var cursor *mongo.Cursor
	var err error
	var res []dto.Item

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.itemsCollection.Find(
			ctx,
			bson.M{"store-id": storeID},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetLimit(int64(count)),
			options.Find().SetSkip(int64(skip)),
		)
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = cursor.All(ctx, &res)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func (s *storage) CreateOrder(order *dto.Order) error {
	logrus.Info("Create Order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.ordersCollection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	logrus.Info(res)

	return nil
}

func (s *storage) UpdateOrder(order *dto.Order) error {
	return nil
}

func (s *storage) DeleteOrder(order *dto.Order) error {
	logrus.Info("Delete Order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := s.ordersCollection.FindOneAndDelete(ctx, bson.M{"_id": order.ID})

	logrus.Info(res)
	return nil
}

func (s *storage) GetOrderList(skip int, count int) (*[]dto.Order, error) {
	logrus.Info("Get Order List")

	var cursor *mongo.Cursor
	var err error
	var res []dto.Order

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.ordersCollection.Find(
			ctx,
			bson.M{},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetLimit(int64(count)),
			options.Find().SetSkip(int64(skip)),
		)
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = cursor.All(ctx, &res)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func (s *storage) CreateStore(store *dto.Store) error {
	logrus.Info("Create Order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.storesCollection.InsertOne(ctx, store)
	if err != nil {
		return err
	}

	logrus.Info(res)

	return nil
}

func (s *storage) UpdateStore(store *dto.Store) error {
	return nil
}

func (s *storage) DeleteStore(store *dto.Store) error {
	logrus.Info("Delete Store")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res := s.storesCollection.FindOneAndDelete(ctx, bson.M{"_id": store.ID})

	logrus.Info(res)
	return nil
}

func (s *storage) GetStoreList(skip int, count int) (*[]dto.Store, error) {
	logrus.Info("Get Store List")

	var cursor *mongo.Cursor
	var err error
	var res []dto.Store

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.storesCollection.Find(
			ctx,
			bson.M{},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetLimit(int64(count)),
			options.Find().SetSkip(int64(skip)),
		)
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = cursor.All(ctx, &res)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}
