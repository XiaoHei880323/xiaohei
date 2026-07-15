package adminResp

type GetUserListAndPageResp struct {
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
	Count    int               `json:"count"`
	UserList []GetUserInfoResp `json:"userList"`
}
type GetUserInfoResp struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	UserNickName string `json:"userNickName"`
	Phone        string `json:"phone"`
	Points       int    `json:"points"`
	Status       int    `json:"status"`
	AddTime      string `json:"addTime"`
}
type GetUserPointsListResp struct {
	Page     int                     `json:"page"`
	PageSize int                     `json:"pageSize"`
	Count    int                     `json:"count"`
	List     []GetUserPointsInfoResp `json:"list"`
}
type GetUserPointsInfoResp struct {
	SourceId int    `json:"sourceId"`
	Points   int    `json:"points"`
	Source   int    `json:"source"`
	CreateAt string `json:"createAt"`
	Notes    string `json:"notes"`
}
