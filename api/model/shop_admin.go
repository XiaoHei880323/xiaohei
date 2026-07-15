package model

type ShopAdmin struct {
	Id       int    `xorm:"not null pk autoincr comment('编号') INT(11)"`
	Username string `xorm:"not null default '' comment('用户名称') index VARCHAR(50)"`
	Password string `xorm:"not null default '' comment('密码') index VARCHAR(100)"`
	Nickname string `xorm:"comment('别名') VARCHAR(50)"`
	Status   int    `xorm:"not null default 0 comment('是否有效  0：无效 1：有效，') index TINYINT(2)"`
}
