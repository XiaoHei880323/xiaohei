package admin

type GetUserListReq struct {
	UserNick string `json:"userNick,optional"`
	Phone    string `json:"phone,optional"`
	Page     int    `json:"page,optional,default=1"`      //当前页
	PageSize int    `json:"pageSize,optional,default=20"` //当前页的条数
}
type UpdateUserPwd struct {
	UserId int `json:"userId"`
}
type UpdateUserPointsReq struct {
	UserId int `json:"userId"`
	Points int `json:"points"`
	Source int `json:"source"` //1:给用户添加积分，2给用户删除积分
}
type GetAdminToUserPointListReq struct {
	UserId   int `json:"userId"`
	Page     int `json:"page,optional,default=1"`      //当前页
	PageSize int `json:"pageSize,optional,default=20"` //当前页的条数
}

// 新增用户
type AddUserReq struct {
	UserName string `json:"userName"`
	RelName  string `json:"relName"`
	Phone    string `json:"phone"`
}

// 修改用户信息
type UpdateUserInfoReq struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName,optional"`
	RelName  string `json:"relName,optional"`
	Phone    string `json:"phone,optional"`
	Status   int    `json:"status,optional"` // 0:正常 1:禁用
}
