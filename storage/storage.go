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

	storeID, err := primitive.ObjectIDFromHex(item.StoreID)
	if err != nil {
		return err
	}

	res, err := s.itemsCollection.InsertOne(ctx, &Item{
		IconUri: item.IconUri,
		Name:    item.Name,
		Price:   item.Price,
		ID:      primitive.NewObjectID(),
		StoreID: storeID,
	})
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

func (s *storage) GetItemListByIdList(idList []string) ([]*Item, error) {
	var err error
	var res []*Item
	var cursor *mongo.Cursor

	oids := make([]primitive.ObjectID, len(idList))
	for i, id := range idList {
		oids[i], err = primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.itemsCollection.Find(ctx, bson.M{"_id": bson.M{"$in": oids}})
		if err != nil {
			return nil, err
		}
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := cursor.All(ctx, &res)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (s *storage) CreateOrder(orderItems map[string]int, location string, storeID string, userID string) error {
	logrus.Info("Create Order")

	order, err := createModelOrder(s, orderItems, location, storeID, userID)
	if err != nil {
		return err
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := s.ordersCollection.InsertOne(ctx, order)
		if err != nil {
			return err
		}

		logrus.Info(res)
	}

	return nil
}

func createModelOrder(s *storage, orderItems map[string]int, location string, storeID string, userID string) (*Order, error) {
	keys := make([]string, len(orderItems))
	i := 0
	for k := range orderItems {
		keys[i] = k
		i++
	}

	items, err := s.GetItemListByIdList(keys)
	if err != nil {
		return nil, err
	}

	// convert model to dto
	orderItemList := make([]*OrderItemData, len(items))
	subtotal := 0
	for i, el := range items {
		orderItemList[i] = &OrderItemData{
			ItemData: &Item{
				IconUri: el.IconUri,
				Name:    el.Name,
				Price:   el.Price,
				StoreID: el.StoreID,
				ID:      el.ID,
			},
			Count: orderItems[el.ID.Hex()],
		}
		subtotal += el.Price * orderItems[el.ID.Hex()]
	}

	storeData, err := s.GetStoreDataByID(storeID)
	if err != nil {
		return nil, err
	}

	userData, err := s.GetUserDataByID(userID)
	if err != nil {
		return nil, err
	}

	orderID := primitive.NewObjectID()

	return &Order{
		ItemList:  orderItemList,
		Location:  location,
		Subtotal:  subtotal,
		StoreData: storeData,
		UserData:  userData,
		ID:        orderID,
		CreatedAt: primitive.Timestamp{T: uint32(time.Now().Unix())},
	}, nil
}

func (s *storage) GetStoreDataByID(id string) (*Store, error) {
	var res Store

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	mongoRes := s.storesCollection.FindOne(ctx, bson.M{"_id": oid})

	err = mongoRes.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *storage) GetUserDataByID(id string) (*User, error) {
	var res User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	mongoRes := s.userCollection.FindOne(ctx, bson.M{"_id": oid})

	err = mongoRes.Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
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

func (s *storage) GetOrderList(skip int, count int, userID string) ([]*dto.Order, error) {
	logrus.Info("Get Order List")

	var cursor *mongo.Cursor
	var err error
	var res []*Order

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err = s.ordersCollection.Find(
			ctx,
			bson.M{"user-data._id": id},
			options.Find().SetSort(bson.M{"created-at": -1}),
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

	// convert model to dto

	return modelOrderListToDtoOrderList(res), nil
}

func modelOrderListToDtoOrderList(orderList []*Order) []*dto.Order {
	dtoOrder := make([]*dto.Order, len(orderList))
	for i, o := range orderList {
		dtoOrder[i] = &dto.Order{
			ItemList:  make([]*dto.OrderItemData, len(o.ItemList)),
			Location:  o.Location,
			CreatedAt: o.CreatedAt.T,
			Subtotal:  o.Subtotal,
			UserData: &dto.User{
				Username: o.UserData.Username,
				ID:       o.UserData.ID.String(),
			},
			StoreData: &dto.Store{
				Location: o.StoreData.Location,
				Name:     o.StoreData.Name,
				ID:       o.StoreData.ID.String(),
			},
			ID: o.ID.String(),
		}

		for k, item := range o.ItemList {
			dtoOrder[i].ItemList[k] = &dto.OrderItemData{
				ItemData: &dto.Item{
					IconUri: item.ItemData.IconUri,
					Name:    item.ItemData.Name,
					Price:   item.ItemData.Price,
					ID:      item.ItemData.ID.String(),
					StoreID: item.ItemData.StoreID.String(),
				},
				Count: item.Count,
			}
		}
	}
	return dtoOrder
}

func (s *storage) CreateStore(store *dto.Store) error {
	logrus.Info("Create Order")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.storesCollection.InsertOne(ctx, &Store{
		Name:     store.Name,
		Location: store.Location,
		ID:       primitive.NewObjectID(),
	})
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

func (s *storage) GetStoreByID(storeID string) (*Store, error) {
	logrus.Info("Get Store")

	var res Store

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, err
	}

	mongoRes := s.storesCollection.FindOne(ctx, bson.M{"_id": id})

	mongoRes.Decode(&res)

	return &res, nil
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

	user.ID = primitive.NewObjectID()
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
