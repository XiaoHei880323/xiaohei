package model

import "time"

type SyActivityConfig struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ConfigName   string    `xorm:"not null default '' comment('配置名称') VARCHAR(100)"`
	ConfigImage  string    `xorm:"not null default '' comment('配置图片') VARCHAR(500)"`
	StartTime    time.Time `xorm:"not null default '2000-01-01 00:00:00' comment('开始时间') DATETIME"`
	EndTime      time.Time `xorm:"not null default '2000-01-01 00:00:00' comment('结束时间') DATETIME"`
	IsDefault    int       `xorm:"not null default 0 comment('是否默认 0:否 1:是') TINYINT(3)"`
	ActivityType int       `xorm:"not null default 1 comment('活动类型 1:签到活动 2:秒杀活动 3:商品活动 4:景点活动') TINYINT(3)"`
	Status       int       `xorm:"not null default 0 comment('状态 0:下线 1:上线') TINYINT(3)"`
	AddTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid       int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid    int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete     int       `xorm:"not null default 0 comment('是否删除 0:未删除 1:已删除') TINYINT(3)"`
}
