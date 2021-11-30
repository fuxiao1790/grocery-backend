package dto

type User struct {
	Username string `bson:"username" json:"username"`
	ID       string `bson:"_id" json:"_id"`
}

type RegisterReq struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type RegisterRes struct {
	Error Err `bson:"error" json:"error"`
}

type LoginReq struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type LoginRes struct {
	UserID   string `bson:"userid" json:"userid"`
	Username string `bson:"username" json:"username"`
	Error    Err    `bson:"error" json:"error"`
}
