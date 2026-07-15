package goodReq

type GoodAddReq struct {
	GoodName  string `json:"goodName"`
	GoodImage string `json:"goodImage,optional"`
	GoodPrice string `json:"goodPrice"`
	GoodDesc  string `json:"goodDesc,optional"`
}
type GoodListReq struct {
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
	GoodName string `json:"goodName,optional"`
}
type GoodUpdateReq struct {
	GoodId    int    `json:"goodId"`
	GoodName  string `json:"goodName"`
	GoodImage string `json:"goodImage,optional"`
	GoodPrice string `json:"goodPrice"`
	GoodDesc  string `json:"goodDesc,optional"`
}
type GoodDeleteReq struct {
	GoodId int `json:"goodId"`
}
type GoodUpdateStatusReq struct {
	GoodId int `json:"goodId"`
	Status int `json:"status"` // 0:下架 1:上架
}
