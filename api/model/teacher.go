package model

import "time"

type Teacher struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	TeacherNo  string    `xorm:"not null unique comment('教师工号') VARCHAR(50)"`
	Name       string    `xorm:"not null comment('教师姓名') VARCHAR(100)"`
	Gender     int       `xorm:"not null default 0 comment('性别 0:未知 1:男 2:女') TINYINT(3)"`
	Phone      string    `xorm:"not null default '' comment('手机号') VARCHAR(20)"`
	Email      string    `xorm:"not null default '' comment('邮箱') VARCHAR(100)"`
	Title      string    `xorm:"not null default '' comment('职称') VARCHAR(100)"`
	Department string    `xorm:"not null default '' comment('所属部门') VARCHAR(100)"`
	Password   string    `xorm:"not null comment('密码') VARCHAR(255)"`
	Status     int       `xorm:"not null default 1 comment('状态 0:禁用 1:正常') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
	AddUid     int       `xorm:"not null default 0 comment('创建人') INT(11)"`
	UpdateUid  int       `xorm:"not null default 0 comment('更新人') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除 0:否 1:是') TINYINT(3)"`
}
