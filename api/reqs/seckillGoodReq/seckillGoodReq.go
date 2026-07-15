package seckillGoodReq

type SeckillGoodListReq struct {
	ActivityId int    `json:"activityId"`
	GoodName   string `json:"goodName,optional"`
	Page       int    `json:"page,optional,default=1"`
	PageSize   int    `json:"pageSize,optional,default=20"`
}

type SeckillGoodAddReq struct {
	ActivityId   int    `json:"activityId"`
	GoodId       int    `json:"goodId"`
	SeckillPrice string `json:"seckillPrice"`
}

type SeckillGoodUpdateReq struct {
	Id           int    `json:"id"`
	SeckillPrice string `json:"seckillPrice"`
}

type SeckillGoodDeleteReq struct {
	Id int `json:"id"`
}

type SeckillGoodBatchUpdateReq struct {
	Ids          []int  `json:"ids"`
	SeckillPrice string `json:"seckillPrice"`
}

type SeckillGoodBatchDeleteReq struct {
	Ids []int `json:"ids"`
}
