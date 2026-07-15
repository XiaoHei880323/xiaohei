package courseMainResp

type CourseMainInfo struct {
	Id              int    `json:"id"`
	StudentId       int    `json:"studentId"`
	StudentName     string `json:"studentName"`
	TeacherId       int    `json:"teacherId"`
	TeacherName     string `json:"teacherName"`
	CourseStartTime string `json:"courseStartTime"`
	CourseEndTime   string `json:"courseEndTime"`
	MeetingLink     string `json:"meetingLink"`
	Subject         string `json:"subject"`
	Status          int    `json:"status"`
	StudentEntered  int    `json:"studentEntered"`
	TeacherEntered  int    `json:"teacherEntered"`
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updateTime"`
}

type CourseMainListResp struct {
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Count    int              `json:"count"`
	List     []CourseMainInfo `json:"list"`
}
