package resp

type UserInfoData struct {
	UserID   int    `json:"user_id"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
}
