package activityConfigResp

// --- 活动配置列表 ---

type ActivityConfigList struct {
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
	Count    int                  `json:"count"`
	List     []ActivityConfigInfo `json:"list"`
}

type ActivityConfigInfo struct {
	ConfigId     int    `json:"configId"`
	ConfigName   string `json:"configName"`
	ConfigImage  string `json:"configImage"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	IsDefault    int    `json:"isDefault"`    // 0:否 1:是
	ActivityType int    `json:"activityType"` // 1:签到 2:秒杀 3:商品 4:景点
	Status       int    `json:"status"`       // 0:下线 1:上线
	AddTime      string `json:"addTime"`
	AddUser      string `json:"addUser"`
}

// --- 活动配置项列表 ---

type ActivityConfigItemList struct {
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
	Count    int                      `json:"count"`
	List     []ActivityConfigItemInfo `json:"list"`
}

type ActivityConfigItemInfo struct {
	ItemId     int    `json:"itemId"`
	ConfigId   int    `json:"configId"`
	ActivityId int    `json:"activityId"`
	Sort       int    `json:"sort"`
	AddTime    string `json:"addTime"`
	AddUser    string `json:"addUser"`
}
