package courseEvaluationReq

type CourseEvaluationListReq struct {
	CourseId int    `json:"courseId,optional"`
	EvalType *int   `json:"evalType,optional"`
	Keyword  string `json:"keyword,optional"`
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
}

type CourseEvaluationDetailReq struct {
	Id int `json:"id"`
}

type CourseEvaluationAddReq struct {
	CourseId int    `json:"courseId"`
	EvalType int    `json:"evalType"`
	Content  string `json:"content"`
	Rating   int    `json:"rating,optional"`
}

type CourseEvaluationUpdateReq struct {
	Id       int    `json:"id"`
	CourseId int    `json:"courseId"`
	EvalType int    `json:"evalType"`
	Content  string `json:"content"`
	Rating   int    `json:"rating,optional"`
}

type CourseEvaluationDeleteReq struct {
	Id int `json:"id"`
}
