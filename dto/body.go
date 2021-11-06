package dto

type GetItemListReq struct {
	Skip  int
	Count int
}

type GetItemListRes struct {
	Items *[]Item
	Error error
}
