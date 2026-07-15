package courseMainReq

type CourseMainListReq struct {
	StudentName string `json:"studentName,optional"`
	TeacherName string `json:"teacherName,optional"`
	StudentId   int    `json:"studentId,optional"`
	TeacherId   int    `json:"teacherId,optional"`
	Status      *int   `json:"status,optional"`
	DateBegin   string `json:"dateBegin,optional"`
	DateEnd     string `json:"dateEnd,optional"`
	Page        int    `json:"page,optional,default=1"`
	PageSize    int    `json:"pageSize,optional,default=20"`
}

type CourseMainDetailReq struct {
	Id int `json:"id"`
}

type CourseMainAddReq struct {
	StudentId       int    `json:"studentId"`
	TeacherId       int    `json:"teacherId"`
	CourseStartTime string `json:"courseStartTime"`
	CourseEndTime   string `json:"courseEndTime"`
	MeetingLink     string `json:"meetingLink,optional"`
	Subject         string `json:"subject,optional"`
	Status          *int   `json:"status,optional"`
}

type CourseMainUpdateReq struct {
	Id              int    `json:"id"`
	StudentId       int    `json:"studentId"`
	TeacherId       int    `json:"teacherId"`
	CourseStartTime string `json:"courseStartTime"`
	CourseEndTime   string `json:"courseEndTime"`
	MeetingLink     string `json:"meetingLink,optional"`
	Subject         string `json:"subject,optional"`
	Status          int    `json:"status,optional"`
}

type CourseMainDeleteReq struct {
	Id int `json:"id"`
}
