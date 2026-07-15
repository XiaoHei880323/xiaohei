package pointsRoleReq

type PointsRoleListReq struct {
	Page                int    `json:"page,optional,default=1"`
	PageSize            int    `json:"pageSize,optional,default=20"`
	GoodName            string `json:"goodName,optional"`
	PointsStart         int    `json:"pointsStart,optional"`
	PointsEnd           int    `json:"pointsEnd,optional"`
	PointsRoleStartTime string `json:"pointsRoleStartTime,optional"`
	PointsRoleEndTime   string `json:"pointsRoleEndTime,optional"`
}

type PointsRoleAddReq struct {
	Points          int    `json:"points"`
	GoodId          int    `json:"goodId"`
	PointsStartTime string `json:"pointsStartTime,optional"`
	PointsEndTime   string `json:"pointsEndTime,optional"`
}

type PointsRoleUpdateReq struct {
	PointId         int    `json:"pointId"`
	Points          int    `json:"points"`
	GoodId          int    `json:"goodId"`
	PointsStartTime string `json:"pointsStartTime,optional"`
	PointsEndTime   string `json:"pointsEndTime,optional"`
}
