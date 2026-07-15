package courseMediaResp

type CourseMediaInfo struct {
	Id         int    `json:"id"`
	CourseId   int    `json:"courseId"`
	MediaType  int    `json:"mediaType"`
	MediaUrl   string `json:"mediaUrl"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type CourseMediaListResp struct {
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Count    int              `json:"count"`
	List     []CourseMediaInfo `json:"list"`
}
