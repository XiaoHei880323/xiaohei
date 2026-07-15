package userResp

// ---- 首页配置 ----

type HomeConfigItem struct {
	ConfigId     int    `json:"configId"`
	ConfigName   string `json:"configName"`
	ConfigImage  string `json:"configImage"`
	ActivityId   int    `json:"activityId"`
	ActivityType int    `json:"activityType"` // 1:签到活动 2:秒杀活动
	Sort         int    `json:"sort"`
}

type HomeConfigResp struct {
	Count int              `json:"count"`
	List  []HomeConfigItem `json:"list"`
}

// ---- 首页数据（活动/商品/景点）----

type ActivityItem struct {
	ActivityId    int    `json:"activityId"`
	ActivityName  string `json:"activityName"`
	ActivityImage string `json:"activityImage"`
	ActivityType  int    `json:"activityType"` // 1:签到 2:秒杀
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	Points        int    `json:"points"`
	IsHot         bool   `json:"isHot"`
}

type GoodItem struct {
	GoodId    int    `json:"goodId"`
	GoodName  string `json:"goodName"`
	GoodImage string `json:"goodImage"`
	GoodPrice string `json:"goodPrice"`
	IsHot     bool   `json:"isHot"`
}

type ScenicItem struct {
	SpotId      int    `json:"spotId"`
	SpotName    string `json:"spotName"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	TicketPrice string `json:"ticketPrice"`
	IsHot       bool   `json:"isHot"`
}

type HomeDataResp struct {
	Activities []ActivityItem `json:"activities"`
	Goods      []GoodItem     `json:"goods"`
	Scenics    []ScenicItem   `json:"scenics"`
}

// ---- 活动配置 ----

type UserActivityConfigItem struct {
	ConfigId     int                          `json:"configId"`
	ConfigName   string                       `json:"configName"`
	ConfigImage  string                       `json:"configImage"`
	StartTime    string                       `json:"startTime"`
	EndTime      string                       `json:"endTime"`
	IsDefault    int                          `json:"isDefault"`
	ActivityType int                          `json:"activityType"` // 1:签到 2:秒杀 3:商品 4:景点
	Items        []UserActivityConfigItemData `json:"items"`
}

type UserActivityConfigItemData struct {
	ItemId     int `json:"itemId"`
	ActivityId int `json:"activityId"`
	Sort       int `json:"sort"`
}

type UserActivityConfigResp struct {
	List []UserActivityConfigItem `json:"list"`
}

// ---- 公告 ----

type NoticeItem struct {
	NoticeId   int    `json:"noticeId"`
	NoticeName string `json:"noticeName"`
}

type NoticeListResp struct {
	Count int          `json:"count"`
	List  []NoticeItem `json:"list"`
}

type NoticeDetailResp struct {
	NoticeId      int    `json:"noticeId"`
	NoticeName    string `json:"noticeName"`
	PublishTime   string `json:"publishTime"`
	PublishUser   string `json:"publishUser"`
	NoticeContent string `json:"noticeContent"`
}
