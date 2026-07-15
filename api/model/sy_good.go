package model

import "time"

type SyGood struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	GoodName   string    `xorm:"not null default '' comment('商品名') VARCHAR(100)"`
	GoodImg    string    `xorm:"not null default '' comment('商品图片') VARCHAR(255)"`
	GoodPrice  string    `xorm:"not null default 0.00 comment('商品价格') DECIMAL(10,2)"`
	GoodDesc   string    `xorm:"comment('商品描述，富文本') LONGTEXT"`
	Status     int       `xorm:"not null default 1 comment('状态 0:下架 1:上架') TINYINT(3)"`
	AddTime    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid     int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid  int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete   int       `xorm:"not null default 0 comment('是否删除') TINYINT(3)"`
}
