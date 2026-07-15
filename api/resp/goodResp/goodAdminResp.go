package goodResp

type GetGoodListResp struct {
	Count    int               `json:"count"`
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
	GoodList []GetGoodInfoResp `json:"goodList"`
}

type GetGoodInfoResp struct {
	GoodId    int    `json:"goodId"`
	GoodName  string `json:"goodName"`
	GoodImg   string `json:"goodImg"`
	GoodPrice string `json:"goodPrice"`
	GoodDesc  string `json:"goodDesc"`
	Status    int    `json:"status"`
}
