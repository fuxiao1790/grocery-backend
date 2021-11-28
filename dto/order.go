package dto

type Order struct {
	Items    map[string]int `bson:"items" json:"items"`
	Location string         `bson:"location" json:"location"`
	UserID   string         `bson:"user-id" json:"user-id"`
	StoreID  string         `bson:"store-id" json:"store-id"`
	ID       string         `bson:"_id" json:"_id"`
}

type CreateOrderReq struct {
	Items    map[string]int `bson:"items" json:"items"`
	Location string         `bson:"location" json:"location"`
	UserID   string         `bson:"user-id" json:"user-id"`
	StoreID  string         `bson:"store-id" json:"store-id"`
}

type CreateOrderRes struct {
	Error Err `bson:"error" json:"error"`
}

type GetOrderListReq struct {
	Skip   int    `bson:"skip" json:"skip"`
	Count  int    `bson:"count" json:"count"`
	UserID string `bson:"user-id" json:"user-id"`
}

type GetOrderListRes struct {
	Orders []*Order `bson:"orders" json:"orders"`
	Error  Err      `bson:"error" json:"error"`
}
