package model

import "time"

type CourseErrorCollection struct {
	Id             int       `xorm:"not null pk autoincr INT(11)"`
	CourseId       int       `xorm:"not null default 0 comment('课程ID') INT(11)"`
	Question       string    `xorm:"not null comment('错题内容') TEXT"`
	CorrectAnswer  string    `xorm:"not null comment('正确答案') TEXT"`
	StudentAnswer  string    `xorm:"not null comment('学生答案') TEXT"`
	Analysis       string    `xorm:"not null comment('解析') TEXT"`
	KnowledgePoint string    `xorm:"not null default '' comment('知识点') VARCHAR(500)"`
	CreateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	AddUid         int       `xorm:"not null default 0 comment('创建人') INT(11)"`
	UpdateUid      int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	IsDelete       int       `xorm:"not null default 0 comment('是否删除 0:否 1:是') TINYINT(3)"`
}
