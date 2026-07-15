package model

import "time"

type SyNotice struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	NoticeName     string    `xorm:"not null default '' comment('公告名称/标题') VARCHAR(200)"`
	NoticeContent  string    `xorm:"comment('公告内容，富文本') LONGTEXT"`
	AddUid         int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	PublishUid     int       `xorm:"not null default 0 comment('发布人ID') INT(11)"`
	PublishTime    time.Time `xorm:"comment('发布时间') TIMESTAMP"`
	AddTime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	UpdateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid      int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete       int       `xorm:"not null default 0 comment('是否删除 0:未删除 1:已删除') TINYINT(3)"`
	NoticeStatus   int       `xorm:"not null default 0 comment('发布状态 0:不发布 1:发布') TINYINT(3)"`
}
