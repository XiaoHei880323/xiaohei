package model

import "time"

type CourseMain struct {
	Id              int       `xorm:"not null pk autoincr INT(11)"`
	StudentId       int       `xorm:"not null default 0 comment('学生ID') INT(11)"`
	TeacherId       int       `xorm:"not null default 0 comment('老师ID') INT(11)"`
	CourseStartTime time.Time `xorm:"not null comment('上课开始时间') TIMESTAMP"`
	CourseEndTime   time.Time `xorm:"not null comment('上课结束时间') TIMESTAMP"`
	MeetingLink     string    `xorm:"not null default '' comment('上课会议链接') VARCHAR(512)"`
	Subject         string    `xorm:"not null default '' comment('课程主题') VARCHAR(200)"`
	Status          int       `xorm:"not null default 0 comment('状态 0:待上课 1:上课中 2:已完成 3:已取消') TINYINT(3)"`
	StudentEntered  int       `xorm:"not null default 0 comment('学生是否进入 0:否 1:是') TINYINT(3)"`
	TeacherEntered  int       `xorm:"not null default 0 comment('老师是否进入 0:否 1:是') TINYINT(3)"`
	CreateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	AddUid          int       `xorm:"not null default 0 comment('创建人') INT(11)"`
	UpdateUid       int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	IsDelete        int       `xorm:"not null default 0 comment('是否删除 0:否 1:是') TINYINT(3)"`
}
