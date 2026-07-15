package model

import (
	"time"
)

type SyPointsRole struct {
	Id                int       `xorm:"not null pk autoincr INT(11)"`
	ExchangePoints    int       `xorm:"not null default 0 comment('兑换积分数') INT(11)"`
	GoodId            int       `xorm:"not null default 0 comment('可兑换的商品') INT(11)"`
	ExchangeStartTime time.Time `xorm:"comment('兑换时间') TIMESTAMP"`
	ExchangeEndTime   time.Time `xorm:"comment('兑换结束时间') TIMESTAMP"`
	CreateTime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid            int       `xorm:"not null default 0 comment('添加人id，sy_admin.id') INT(11)"`
	UpdateTime        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间，') TIMESTAMP"`
	UpdateUid         int       `xorm:"not null default 0 comment('修改用户时间') INT(11)"`
	IsDeleted         int       `xorm:"not null default 0 comment('是否删除，0：未删除，1：删除') TINYINT(3)"`
}
