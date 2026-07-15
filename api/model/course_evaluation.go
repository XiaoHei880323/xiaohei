package model

import "time"

type CourseEvaluation struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	CourseId   int       `xorm:"not null default 0 comment('课程ID') INT(11)"`
	EvalType   int       `xorm:"not null default 0 comment('类型 0:学生对老师 1:家长对老师 3:老师对学生') TINYINT(3)"`
	Content    string    `xorm:"not null comment('评价内容') TEXT"`
	Rating     int       `xorm:"not null default 5 comment('评分 1-5') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	AddUid     int       `xorm:"not null default 0 comment('创建人') INT(11)"`
	UpdateUid  int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除 0:否 1:是') TINYINT(3)"`
}
