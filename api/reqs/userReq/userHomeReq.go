package userReq

type HomeConfigReq struct {
	Page     int `json:"page,optional,default=1"`
	PageSize int `json:"pageSize,optional,default=10"`
}

type HomeDataReq struct {
	Page     int `json:"page,optional,default=1"`
	PageSize int `json:"pageSize,optional,default=10"`
}

type NoticeListReq struct {
	Page     int `json:"page,optional,default=1"`
	PageSize int `json:"pageSize,optional,default=10"`
}

type NoticeDetailReq struct {
	NoticeId int `json:"noticeId"`
}

type UserActivityConfigReq struct{}

