package activityResp

type GetActivityList struct {
	Page         int                       `json:"page"`
	PageSize     int                       `json:"pageSize"`
	Count        int                       `json:"count"`
	ActivityList []GetActivityListInfoResp `json:"activityList"`
}

type GetActivityListInfoResp struct {
	ActivityId          int    `json:"activityId"`
	ActivityName        string `json:"activityName"`
	ActivityImage       string `json:"activityImage"`
	ActivityText        string `json:"activityText"`
	ActivityStartTime   string `json:"activityStartTime"`
	ActivityEndTime     string `json:"activityEndTime"`
	ActivityPreviewTime string `json:"activityPreviewTime"`
	ActivityPoints      int    `json:"activityPoints"`
	ActivityAddTime     string `json:"activityAddTime"`
	ActivityAddUser     string `json:"activityAddUser"`
}

type GetActivityInfoResp struct {
	ActivityId          int    `json:"activityId"`
	ActivityName        string `json:"activityName"`
	ActivityImage       string `json:"activityImage"`
	ActivityText        string `json:"activityText"`
	ActivityStartTime   string `json:"activityStartTime"`
	ActivityEndTime     string `json:"activityEndTime"`
	ActivityPreviewTime string `json:"activityPreviewTime"`
	ActivityPoints      string `json:"activityPoints"`
	ActivityAddTime     string `json:"activityAddTime"`
	ActivityAddUser     string `json:"activityAddUser"`
}
