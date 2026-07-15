package model

import (
	"time"
)

type SyUserPoints struct {
	Id         int       `xorm:"not null pk autoincr INT(11)"`
	UserId     int       `xorm:"not null default 0 comment('用户id sy_user.id') INT(11)"`
	SourceId   int       `xorm:"not null default 0 comment('对应的id ，sy_activity.id or sy_points_role.id') INT(11)"`
	Points     int       `xorm:"not null default 0 comment('对应的积分数') INT(11)"`
	Source     int       `xorm:"not null default 0 comment('积分的来源,1:用户参加活动，2用户兑换积分 3：管理员添加用户积分；4：管理管减少用户积分') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('用户积分变化时间') TIMESTAMP"`
	Notes      string    `xorm:"comment('备注') VARCHAR(255)"`
	IsDelete   int       `xorm:"not null default 0 comment('0；正常，1删除') TINYINT(3)"`
}
