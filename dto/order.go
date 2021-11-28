package dto

type OrderItem struct {
	IconUri string `bson:"icon-uri" json:"icon-uri"`
	Name    string `bson:"name" json:"name"`
	Price   string `bson:"price" json:"price"`
	Count   int    `bson:"Count" json:"Count"`
	ID      string `bson:"_id" json:"_id"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type Order struct {
	Items    map[*Item]int `bson:"items" json:"items"`
	Location string        `bson:"location" json:"location"`
	UserID   string        `bson:"user-id" json:"user-id"`
	ID       string        `bson:"_id" json:"_id"`
	StoreID  string        `bson:"store-id" json:"store-id"`
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
