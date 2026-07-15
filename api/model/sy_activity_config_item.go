package model

import "time"

type SyActivityConfigItem struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	ConfigId   int       `xorm:"not null default 0 comment('活动配置ID') INT(11)"`
	ActivityId int       `xorm:"not null default 0 comment('关联活动/商品/景点ID') INT(11)"`
	Sort       int       `xorm:"not null default 0 comment('排序') INT(11)"`
	AddTime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid     int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid  int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除 0:未删除 1:已删除') TINYINT(3)"`
}
