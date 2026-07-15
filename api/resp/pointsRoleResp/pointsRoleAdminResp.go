package pointsRoleResp

type GetPointsRoleListResp struct {
	Count          int                     `json:"count"`
	PageSize       int                     `json:"pageSize"`
	Page           int                     `json:"page"`
	PointsRoleList []GetPointsRoleInfoResp `json:"pointsRoleList"`
}
type GetPointsRoleInfoResp struct {
	PointId        int    `json:"pointId"`
	Point          int    `json:"point"`
	GoodId         int    `json:"goodId"`
	GoodName       string `json:"goodName"`
	GoodImage      string `json:"goodImage"`
	PointStartTime string `json:"pointStartTime"`
	PointEndTime   string `json:"pointEndTime"`
	CreateTime     string `json:"createTime"`
}
