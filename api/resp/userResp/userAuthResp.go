package userResp

type UserLoginResp struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Phone    string `json:"phone"`
	HeadImg  string `json:"headImg"`
	Points   int    `json:"points"`
	Token    string `json:"token"`
}
