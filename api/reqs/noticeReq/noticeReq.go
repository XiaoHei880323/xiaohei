package noticeReq

type NoticeListReq struct {
	NoticeName string `json:"noticeName,optional"`
	Page       int    `json:"page,optional,default=1"`
	PageSize   int    `json:"pageSize,optional,default=20"`
}

type NoticeAddReq struct {
	NoticeName    string `json:"noticeName"`
	NoticeContent string `json:"noticeContent,optional"`
	PublishUid    int    `json:"publishUid,optional"`
	NoticeStatus  int    `json:"noticeStatus,optional"` // 0:不发布 1:发布
}

type NoticeUpdateReq struct {
	NoticeId      int    `json:"noticeId"`
	NoticeName    string `json:"noticeName"`
	NoticeContent string `json:"noticeContent,optional"`
	PublishUid    int    `json:"publishUid,optional"`
	NoticeStatus  int    `json:"noticeStatus,optional"` // 0:不发布 1:发布
}

type NoticeDeleteReq struct {
	NoticeId int `json:"noticeId"`
}

type NoticeStatusReq struct {
	NoticeId     int `json:"noticeId"`
	NoticeStatus int `json:"noticeStatus"` // 0:下线 1:发布
}
