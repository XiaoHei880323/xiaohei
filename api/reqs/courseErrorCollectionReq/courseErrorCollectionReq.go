package courseErrorCollectionReq

type CourseErrorCollectionListReq struct {
	CourseId int    `json:"courseId,optional"`
	Keyword  string `json:"keyword,optional"`
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
}

type CourseErrorCollectionDetailReq struct {
	Id int `json:"id"`
}

type CourseErrorCollectionAddReq struct {
	CourseId       int    `json:"courseId"`
	Question       string `json:"question"`
	CorrectAnswer  string `json:"correctAnswer"`
	StudentAnswer  string `json:"studentAnswer,optional"`
	Analysis       string `json:"analysis,optional"`
	KnowledgePoint string `json:"knowledgePoint,optional"`
}

type CourseErrorCollectionUpdateReq struct {
	Id             int    `json:"id"`
	CourseId       int    `json:"courseId"`
	Question       string `json:"question"`
	CorrectAnswer  string `json:"correctAnswer"`
	StudentAnswer  string `json:"studentAnswer,optional"`
	Analysis       string `json:"analysis,optional"`
	KnowledgePoint string `json:"knowledgePoint,optional"`
}

type CourseErrorCollectionDeleteReq struct {
	Id int `json:"id"`
}
