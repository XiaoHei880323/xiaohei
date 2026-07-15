package model

import "time"

type CourseMedia struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	CourseId   int       `xorm:"not null default 0 comment('课程ID') INT(11)"`
	MediaType  int       `xorm:"not null default 0 comment('类型 0:学生上传 1:老师上传 3:上课录音') TINYINT(3)"`
	MediaUrl   string    `xorm:"not null default '' comment('资源地址') VARCHAR(1024)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	AddUid     int       `xorm:"not null default 0 comment('创建人') INT(11)"`
	UpdateUid  int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除 0:否 1:是') TINYINT(3)"`
}
