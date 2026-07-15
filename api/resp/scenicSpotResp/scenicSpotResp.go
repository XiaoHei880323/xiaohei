package scenicSpotResp

type GetScenicSpotList struct {
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Count    int              `json:"count"`
	List     []ScenicSpotInfo `json:"list"`
}

type ScenicSpotInfo struct {
	Id          int    `json:"id"`
	SpotName    string `json:"spotName"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	TicketPrice string `json:"ticketPrice"`
	Description string `json:"description"`
	AddTime     string `json:"addTime"`
	Status      int    `json:"status"`
}
