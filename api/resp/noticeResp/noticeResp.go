package noticeResp

type NoticeList struct {
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
	Count    int          `json:"count"`
	List     []NoticeInfo `json:"list"`
}

type NoticeInfo struct {
	NoticeId      int    `json:"noticeId"`
	NoticeName    string `json:"noticeName"`
	PublishTime   string `json:"publishTime"`
	NoticeContent string `json:"noticeContent"`
	NoticeStatus  int    `json:"noticeStatus"` // 0:不发布 1:发布
	AddUser       string `json:"addUser"`
	PublishUser   string `json:"publishUser"`
	AddTime       string `json:"addTime"`
}
