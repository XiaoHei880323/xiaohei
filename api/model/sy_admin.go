package model

import (
	"time"
)

type SyAdmin struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	UserName   string    `xorm:"not null default '无' comment('用户名') VARCHAR(50)"`
	RelName    string    `xorm:"not null default '无' comment('昵称') VARCHAR(100)"`
	Phone      int64     `xorm:"comment('手机号码') index BIGINT(20)"`
	Pwd        string    `xorm:"not null default '' comment('密码') VARCHAR(255)"`
	Status     int       `xorm:"not null default 0 comment('状态；0：有效，1：无效') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('新增时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	AddAid     int       `xorm:"not null default 0 comment('新增人') INT(11)"`
	UpdateAid  int       `xorm:"not null default 0 comment('修改人') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除，0：未删除 1：已删除') TINYINT(3)"`
}
