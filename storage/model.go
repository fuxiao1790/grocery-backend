package storage

type Item struct {
	IconUri string `bson:"icon-uri" json:"icon-uri"`
	Name    string `bson:"name" json:"name"`
	Price   string `bson:"price" json:"price"`
	ID      string `bson:"_id" json:"_id"`
	StoreID string `bson:"store-id" json:"store-id"`
}

type Order struct {
	Items    map[*Item]int `bson:"items" json:"items"`
	Location string        `bson:"location" json:"location"`
	ID       string        `bson:"_id" json:"_id"`
	StoreID  string        `bson:"store-id" json:"store-id"`
}

type Store struct {
	Location string `bson:"location" json:"location"`
	ID       string `bson:"_id" json:"_id"`
	Name     string `bson:"name" json:"name"`
}
