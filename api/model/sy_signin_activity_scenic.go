package model

import "time"

type SySigninActivityScenic struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	ActivityId  int       `xorm:"not null default 0 comment('签到活动ID') INT(11)"`
	ScenicId    int       `xorm:"not null default 0 comment('景点ID') INT(11)"`
	SignPoints  int       `xorm:"not null default 0 comment('签到可获积分') INT(11)"`
	QrCodeUrl   string    `xorm:"not null default '' comment('打卡二维码地址') VARCHAR(500)"`
	Status      int       `xorm:"not null default 1 comment('状态 0:禁用 1:启用') TINYINT(3)"`
	AddTime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid      int       `xorm:"not null default 0 comment('添加人ID') INT(11)"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid   int       `xorm:"not null default 0 comment('修改人ID') INT(11)"`
	IsDelete    int       `xorm:"not null default 0 comment('是否删除') TINYINT(3)"`
}
