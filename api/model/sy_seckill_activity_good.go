package model

import "time"

type SySeckillActivityGood struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ActivityId   int       `xorm:"not null default 0 comment('秒杀活动ID') INT(11)"`
	GoodId       int       `xorm:"not null default 0 comment('商品ID') INT(11)"`
	SeckillPrice string    `xorm:"not null default 0.00 comment('秒杀价格') DECIMAL(10,2)"`
	AddTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid       int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid    int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete     int       `xorm:"not null default 0 comment('是否删除 0未删除 1已删除') TINYINT(3)"`
}
