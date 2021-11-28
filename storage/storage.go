package storage

import (
	"context"
	"errors"
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
const USER_COLLECTION_NAME = "USERS"

type storage struct {
	client           *mongo.Client
	itemsCollection  *mongo.Collection
	ordersCollection *mongo.Collection
	storesCollection *mongo.Collection
	userCollection   *mongo.Collection
}

type Config struct {
	Uri  string
	Name string
}

var INVALID_ID = errors.New("inavlid id")

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
	re.userCollection = re.client.Database(config.Name).Collection(USER_COLLECTION_NAME)

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

func (s *storage) ValidateID(id string) bool {
	return primitive.IsValidObjectID(id)
}

func (s *storage) GetItemList(skip int, count int, storeID string) ([]*dto.Item, error) {
	logrus.Info("Get Item List")

	var cursor *mongo.Cursor
	var err error
	var res []*dto.Item

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		storeOjbectID, err := primitive.ObjectIDFromHex(storeID)
		if err != nil {
			return nil, INVALID_ID
		}

		cursor, err = s.itemsCollection.Find(
			ctx,
			bson.M{"store-id": storeOjbectID},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetSkip(int64(skip)),
			options.Find().SetLimit(int64(count)),
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

	logrus.Debugf("skip: %d, count: %d, actual: %d", skip, count, len(res))

	return res, nil
}

func (s *storage) CreateOrder(order *dto.Order) error {
	logrus.Info("Create Order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert data from dto order to model order to store in db
	mOrder := &Order{
		Items:    order.Items,
		Location: order.Location,
	}

	{
		id, err := primitive.ObjectIDFromHex(order.StoreID)
		if err != nil {
			return err
		}
		mOrder.StoreID = id
	}
	{
		id, err := primitive.ObjectIDFromHex(order.UserID)
		if err != nil {
			return err
		}
		mOrder.UserID = id
	}

	mOrder.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	res, err := s.ordersCollection.InsertOne(ctx, mOrder)
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

func (s *storage) GetOrderList(skip int, count int) ([]*dto.Order, error) {
	logrus.Info("Get Order List")

	var cursor *mongo.Cursor
	var err error
	var res []*dto.Order

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.ordersCollection.Find(
			ctx,
			bson.M{},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetSkip(int64(skip)),
			options.Find().SetLimit(int64(count)),
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

	logrus.Debugf("skip: %d, count: %d, actual: %d", skip, count, len(res))

	return res, nil
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

func (s *storage) GetStoreList(skip int, count int) ([]*dto.Store, error) {
	logrus.Info("Get Store List")

	var cursor *mongo.Cursor
	var err error
	var res []*dto.Store

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.storesCollection.Find(
			ctx,
			bson.M{},
			options.Find().SetSort(bson.M{"_id": -1}),
			options.Find().SetSkip(int64(skip)),
			options.Find().SetLimit(int64(count)),
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

	logrus.Debugf("skip: %d, count: %d, actual: %d", skip, count, len(res))

	return res, nil
}

func (s *storage) CreateUser(user *User) error {
	logrus.Info("Create User")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectIDFromTimestamp(time.Now())
	res, err := s.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	logrus.Debug(res)

	return nil
}

func (s *storage) GetUser(user *User) (*User, error) {
	logrus.Info("Get User")

	var res User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoRes := s.userCollection.FindOne(ctx, bson.M{"username": user.Username})
	if mongoRes.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}
	if mongoRes.Err() != nil {
		return nil, mongoRes.Err()
	}

	mongoRes.Decode(&res)

	logrus.Debug(res)

	return &res, nil
}
