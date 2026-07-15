package activityReq

type ActivityListReq struct {
	ActivityName      string `json:"activityName,optional"`        //活动名长
	ActivityStartTime string `json:"activityStartTime,optional"`   //活动开始时间
	ActivityEndTime   string `json:"activityEndTime,optional"`     //活动结束时间
	Page              int    `json:"page,optional,default=1"`      //当前页
	PageSize          int    `json:"pageSize,optional,default=20"` //当前页的条数
}
type ActivityAddInfoReq struct {
	ActivityName        string `json:"activityName"`
	ActivityImage       string `json:"activityImage,optional"`
	ActivityText        string `json:"activityText,optional"`
	ActivityStartTime   string `json:"activityStartTime,optional"`
	ActivityEndTime     string `json:"activityEndTime,optional"`
	ActivityPreviewTime string `json:"activityPreviewTime,optional"`
	ActivityPoints      int    `json:"activityPoints,optional"`
}
type ActivityUpdateInfoReq struct {
	ActivityId          int    `json:"activityId"`
	ActivityName        string `json:"activityName"`
	ActivityImage       string `json:"activityImage,optional"`
	ActivityText        string `json:"activityText,optional"`
	ActivityStartTime   string `json:"activityStartTime,optional"`
	ActivityEndTime     string `json:"activityEndTime,optional"`
	ActivityPreviewTime string `json:"activityPreviewTime,optional"`
	ActivityPoints      int    `json:"activityPoints,optional"`
}
type ActivityDeleteReq struct {
	ActivityId int `json:"activityId"`
}
