package admin

type UserList struct {
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
	UserName string `json:"userName,optional"`
}

type UserLogin struct {
	Phone string `json:"phone"`
	Pwd   string `json:"pwd"`
}
type UpdateInfoReq struct {
	UserId    int    `json:"userId"`
	OldPwd    string `json:"oldPwd,optional"`
	NewOldPwd string `json:"newOldPwd,optional"`
	UserName  string `json:"userName,optional"`
	RelName   string `json:"relName,optional"`
	Phone     int    `json:"phone,optional"`
}
type GetAdminUserListReq struct {
	UserName string `json:"userName,optional"`
	RelName  string `json:"relName,optional"`
	Phone    string `json:"phone,optional"`
	Page     int    `json:"page,optional,default=1"`
	PageSize int    `json:"pageSize,optional,default=20"`
}
type AddAdminUserReq struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	RelName  string `json:"relName"`
	Phone    string `json:"phone"`
}

type UpdateAdminUserReq struct {
	UserId       int `json:"userId"`
	UpdateUserId int `json:"updateUserId"`
	Type         int `json:"type,optional,default=1"`
}
