package scenicSpotReq

type ScenicSpotListReq struct {
	SpotName string `json:"spotName,optional"`
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
}

type ScenicSpotAddReq struct {
	SpotName    string `json:"spotName"`
	Longitude   string `json:"longitude,optional"`
	Latitude    string `json:"latitude,optional"`
	TicketPrice string `json:"ticketPrice,optional"`
	Description string `json:"description,optional"`
}

type ScenicSpotUpdateReq struct {
	Id          int    `json:"id"`
	SpotName    string `json:"spotName"`
	Longitude   string `json:"longitude,optional"`
	Latitude    string `json:"latitude,optional"`
	TicketPrice string `json:"ticketPrice,optional"`
	Description string `json:"description,optional"`
}

type ScenicSpotDeleteReq struct {
	Id int `json:"id"`
}
type ScenicSpotUpdateStatusReq struct {
	Id     int `json:"id"`
	Status int `json:"status"` // 0:下架 1:上架
}
