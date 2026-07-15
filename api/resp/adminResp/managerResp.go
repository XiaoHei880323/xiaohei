package adminResp

type AdminListResp struct {
	Count    int             `json:"count"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	List     []AdminInfoResp `json:"list"`
}

type AdminInfoResp struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Status   int    `json:"status"`
	CreateAt string `json:"createAt"`
	Token    string `json:"token"`
}
type GetAdminUserListResp struct {
	Page      int             `json:"page"`
	PageSize  int             `json:"pageSize"`
	Count     int             `json:"count"`
	AdminList []AdminInfoResp `json:"adminList"`
}
