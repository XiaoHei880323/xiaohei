package courseErrorCollectionResp

type CourseErrorCollectionInfo struct {
	Id             int    `json:"id"`
	CourseId       int    `json:"courseId"`
	Question       string `json:"question"`
	CorrectAnswer  string `json:"correctAnswer"`
	StudentAnswer  string `json:"studentAnswer"`
	Analysis       string `json:"analysis"`
	KnowledgePoint string `json:"knowledgePoint"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
}

type CourseErrorCollectionListResp struct {
	Page     int                        `json:"page"`
	PageSize int                        `json:"pageSize"`
	Count    int                        `json:"count"`
	List     []CourseErrorCollectionInfo `json:"list"`
}
