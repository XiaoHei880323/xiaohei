package model

import "time"

type SyHomeConfig struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ConfigName   string    `xorm:"not null default '' comment('配置名称') VARCHAR(100)"`
	ConfigImage  string    `xorm:"not null default '' comment('配置图片') VARCHAR(500)"`
	ActivityId   int       `xorm:"not null default 0 comment('活动ID，sy_activity.id') INT(11)"`
	ActivityType int       `xorm:"not null default 1 comment('活动类型 1:签到活动 2:秒杀活动') TINYINT(3)"`
	Sort         int       `xorm:"not null default 0 comment('排序') INT(11)"`
	AddTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid       int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid    int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete     int       `xorm:"not null default 0 comment('是否删除 0:未删除 1:已删除') TINYINT(3)"`
}
