package teacherResp

type TeacherInfo struct {
	Id         int    `json:"id"`
	TeacherNo  string `json:"teacherNo"`
	Name       string `json:"name"`
	Gender     int    `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Title      string `json:"title"`
	Department string `json:"department"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type TeacherListResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Count    int           `json:"count"`
	List     []TeacherInfo `json:"list"`
}
