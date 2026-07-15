package studentResp

type StudentInfo struct {
	Id         int    `json:"id"`
	StudentNo  string `json:"studentNo"`
	Name       string `json:"name"`
	Gender     int    `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	ClassName  string `json:"className"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type StudentListResp struct {
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Count    int           `json:"count"`
	List     []StudentInfo `json:"list"`
}
