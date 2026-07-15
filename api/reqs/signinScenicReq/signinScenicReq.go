package signinScenicReq

type SigninScenicListReq struct {
	ActivityId int `json:"activityId"`
	Page       int `json:"page,optional,default=1"`
	PageSize   int `json:"pageSize,optional,default=50"`
}

type SigninScenicAddReq struct {
	ActivityId int    `json:"activityId"`
	ScenicId   int    `json:"scenicId"`
	SignPoints int    `json:"signPoints"`
	QrCodeUrl  string `json:"qrCodeUrl,optional"`
	Status     int    `json:"status,optional"`
}

type SigninScenicUpdateReq struct {
	Id         int    `json:"id"`
	SignPoints int    `json:"signPoints"`
	QrCodeUrl  string `json:"qrCodeUrl,optional"`
	Status     int    `json:"status,optional"`
}

type SigninScenicDeleteReq struct {
	Id int `json:"id"`
}
