package model

import "time"

type SyScenicSpot struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	SpotName     string    `xorm:"not null default '' comment('景点名称') VARCHAR(100)"`
	Longitude    string    `xorm:"not null default 0.000000 comment('经度') DECIMAL(10,6)"`
	Latitude     string    `xorm:"not null default 0.000000 comment('纬度') DECIMAL(10,6)"`
	TicketPrice  string    `xorm:"not null default 0.00 comment('票价') DECIMAL(10,2)"`
	Description  string    `xorm:"comment('景点描述，富文本') LONGTEXT"`
	AddTime      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid       int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid    int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete     int       `xorm:"not null default 0 comment('是否删除') TINYINT(3)"`
	Status       int       `xorm:"not null default 1 comment('状态 0:下架 1:上架') TINYINT(3)"`
}
