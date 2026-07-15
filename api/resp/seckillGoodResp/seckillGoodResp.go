package seckillGoodResp

type GetSeckillGoodList struct {
	Page     int                `json:"page"`
	PageSize int                `json:"pageSize"`
	Count    int                `json:"count"`
	List     []GetSeckillGoodInfo `json:"list"`
}

type GetSeckillGoodInfo struct {
	Id           int    `json:"id"`
	ActivityId   int    `json:"activityId"`
	GoodId       int    `json:"goodId"`
	GoodName     string `json:"goodName"`
	GoodImg      string `json:"goodImg"`
	GoodPrice    string `json:"goodPrice"`
	SeckillPrice string `json:"seckillPrice"`
	AddTime      string `json:"addTime"`
}
