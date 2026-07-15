package userReq

type UserLoginReq struct {
	Phone string `json:"phone"`
	Pwd   string `json:"pwd"`
}
