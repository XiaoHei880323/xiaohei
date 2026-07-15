package teacherReq

type TeacherListReq struct {
	Keyword    string `json:"keyword,optional"`
	Department string `json:"department,optional"`
	Status     *int   `json:"status,optional"`
	Page       int    `json:"page,optional,default=1"`
	PageSize   int    `json:"pageSize,optional,default=20"`
}

type TeacherDetailReq struct {
	Id int `json:"id"`
}

type TeacherAddReq struct {
	TeacherNo  string `json:"teacherNo"`
	Name       string `json:"name"`
	Gender     int    `json:"gender,optional"`
	Phone      string `json:"phone,optional"`
	Email      string `json:"email,optional"`
	Title      string `json:"title,optional"`
	Department string `json:"department,optional"`
	Password   string `json:"password"`
	Status     *int   `json:"status,optional"`
}

type TeacherUpdateReq struct {
	Id         int    `json:"id"`
	TeacherNo  string `json:"teacherNo"`
	Name       string `json:"name"`
	Gender     int    `json:"gender,optional"`
	Phone      string `json:"phone,optional"`
	Email      string `json:"email,optional"`
	Title      string `json:"title,optional"`
	Department string `json:"department,optional"`
	Password   string `json:"password,optional"`
	Status     int    `json:"status,optional"`
}

type TeacherDeleteReq struct {
	Id int `json:"id"`
}
