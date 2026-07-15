package courseEvaluationResp

type CourseEvaluationInfo struct {
	Id         int    `json:"id"`
	CourseId   int    `json:"courseId"`
	EvalType   int    `json:"evalType"`
	Content    string `json:"content"`
	Rating     int    `json:"rating"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type CourseEvaluationListResp struct {
	Page     int                   `json:"page"`
	PageSize int                   `json:"pageSize"`
	Count    int                   `json:"count"`
	List     []CourseEvaluationInfo `json:"list"`
}
