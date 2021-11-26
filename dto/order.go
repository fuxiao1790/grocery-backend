package dto

type Order struct {
	Items    map[*Item]int `bson:"items" json:"items"`
	Location string        `bson:"location" json:"location"`
	ID       string        `bson:"_id" json:"_id"`
	StoreID  string        `bson:"store-id" json:"store-id"`
}

type CreateOrderReq struct {
	Order Order `bson:"order" json:"order"`
}

type CreateOrderRes struct {
	Error error `bson:"error" json:"error"`
}

type GetOrderListReq struct {
	Skip  int `bson:"skip" json:"skip"`
	Count int `bson:"count" json:"count"`
}

type GetOrderListRes struct {
	Orders []*Order `bson:"orders" json:"orders"`
	Error  error    `bson:"error" json:"error"`
}
