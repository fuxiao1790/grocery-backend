package dto

type Item struct {
	IconUri string `bson:"icon-uri" json:"icon-uri"`
	Name    string `bson:"name" json:"name"`
	Price   int    `bson:"price" json:"price"`
	ID      string `bson:"_id" json:"_id"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type GetItemListReq struct {
	Skip    int    `bson:"skip" json:"skip"`
	Count   int    `bson:"count" json:"count"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type GetItemListRes struct {
	Items []*Item `bson:"items" json:"items"`
	Error Err     `bson:"error" json:"error"`
}

type NewItemReq struct {
	IconUri string `bson:"icon-uri" json:"icon-uri"`
	Name    string `bson:"name" json:"name"`
	Price   int    `bson:"price" json:"price"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type NewItemRes struct {
	Error Err `bson:"error" json:"error"`
}
