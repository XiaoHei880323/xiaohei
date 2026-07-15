package studentReq

type StudentListReq struct {
	Keyword   string `json:"keyword,optional"`
	ClassName string `json:"className,optional"`
	Status    *int   `json:"status,optional"`
	Page      int    `json:"page,optional,default=1"`
	PageSize  int    `json:"pageSize,optional,default=20"`
}

type StudentDetailReq struct {
	Id int `json:"id"`
}

type StudentAddReq struct {
	StudentNo string `json:"studentNo"`
	Name      string `json:"name"`
	Gender    int    `json:"gender,optional"`
	Phone     string `json:"phone,optional"`
	Email     string `json:"email,optional"`
	ClassName string `json:"className,optional"`
	Password  string `json:"password"`
	Status    *int   `json:"status,optional"`
}

type StudentUpdateReq struct {
	Id        int    `json:"id"`
	StudentNo string `json:"studentNo"`
	Name      string `json:"name"`
	Gender    int    `json:"gender,optional"`
	Phone     string `json:"phone,optional"`
	Email     string `json:"email,optional"`
	ClassName string `json:"className,optional"`
	Password  string `json:"password,optional"`
	Status    int    `json:"status,optional"`
}

type StudentDeleteReq struct {
	Id int `json:"id"`
}
