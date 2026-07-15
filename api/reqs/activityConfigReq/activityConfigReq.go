package activityConfigReq

// --- 活动配置 ---

type ActivityConfigListReq struct {
	ConfigName   string `json:"configName,optional"`
	ActivityType int    `json:"activityType,optional"`
	Page         int    `json:"page,optional,default=1"`
	PageSize     int    `json:"pageSize,optional,default=20"`
}

type ActivityConfigAddReq struct {
	ConfigName   string `json:"configName"`
	ConfigImage  string `json:"configImage,optional"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	ActivityType int    `json:"activityType"`
}

type ActivityConfigUpdateReq struct {
	ConfigId     int    `json:"configId"`
	ConfigName   string `json:"configName"`
	ConfigImage  string `json:"configImage,optional"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	ActivityType int    `json:"activityType"`
}

type ActivityConfigDeleteReq struct {
	ConfigId int `json:"configId"`
}

type ActivityConfigSetDefaultReq struct {
	ConfigId int `json:"configId"`
}

type ActivityConfigUpdateStatusReq struct {
	ConfigId int `json:"configId"`
	Status   int `json:"status"` // 0:下线 1:上线
}

// --- 活动配置项 ---

type ActivityConfigItemListReq struct {
	ConfigId int `json:"configId"`
	Page     int `json:"page,optional,default=1"`
	PageSize int `json:"pageSize,optional,default=20"`
}

type ActivityConfigItemAddReq struct {
	ConfigId   int `json:"configId"`
	ActivityId int `json:"activityId"`
	Sort       int `json:"sort,optional"`
}

type ActivityConfigItemUpdateReq struct {
	ItemId     int `json:"itemId"`
	ActivityId int `json:"activityId"`
	Sort       int `json:"sort,optional"`
}

type ActivityConfigItemDeleteReq struct {
	ItemId int `json:"itemId"`
}
