package signinScenicResp

type GetSigninScenicList struct {
	Page        int               `json:"page"`
	PageSize    int               `json:"pageSize"`
	Count       int               `json:"count"`
	TotalPoints int               `json:"totalPoints"`
	List        []SigninScenicInfo `json:"list"`
}

type SigninScenicInfo struct {
	Id         int    `json:"id"`
	ActivityId int    `json:"activityId"`
	ScenicId   int    `json:"scenicId"`
	SpotName   string `json:"spotName"`
	SignPoints int    `json:"signPoints"`
	QrCodeUrl  string `json:"qrCodeUrl"`
	Status     int    `json:"status"`
	AddTime    string `json:"addTime"`
}
