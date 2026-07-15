package courseMediaReq

type CourseMediaListReq struct {
	CourseId  int  `json:"courseId,optional"`
	MediaType *int `json:"mediaType,optional"`
	Page      int  `json:"page,optional,default=1"`
	PageSize  int  `json:"pageSize,optional,default=20"`
}

type CourseMediaDetailReq struct {
	Id int `json:"id"`
}

type CourseMediaAddReq struct {
	CourseId  int    `json:"courseId"`
	MediaType int    `json:"mediaType"`
	MediaUrl  string `json:"mediaUrl"`
}

type CourseMediaUpdateReq struct {
	Id        int    `json:"id"`
	CourseId  int    `json:"courseId"`
	MediaType int    `json:"mediaType"`
	MediaUrl  string `json:"mediaUrl"`
}

type CourseMediaDeleteReq struct {
	Id int `json:"id"`
}
