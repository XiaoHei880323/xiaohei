package model

import (
	"time"
)

type SyActivity struct {
	Id                   int       `xorm:"not null pk autoincr INT(11)"`
	ActivityType         int       `xorm:"not null default 1 comment('活动类型，1：签到活动，2：秒杀活动') TINYINT(3)"`
	ActivityName         string    `xorm:"not null default '标题' comment('活动标题') VARCHAR(50)"`
	ActivityImg          string    `xorm:"default '' comment('活动首页图片') VARCHAR(255)"`
	ActivityText         string    `xorm:"not null comment('正文，富文本存储') TEXT"`
	ActivityStartingTime time.Time `xorm:"comment('活动开始时间') TIMESTAMP"`
	ActivityEndTime      time.Time `xorm:"comment('活动结束时间') TIMESTAMP"`
	ActivityPreviewTime  time.Time `xorm:"comment('活动预告时间') TIMESTAMP"`
	ActivityPoints       int       `xorm:"not null default 0 comment('参加活动可以获得多少积分') INT(11)"`
	AddTime              time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	AddUid               int       `xorm:"not null default 0 comment('添加用户id，sy_admin.id') INT(11)"`
	UpdateTime           time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
	UpdateUid            int       `xorm:"not null default 0 comment('修改用户id，sy_admin.id') INT(11)"`
	IsDelete             int       `xorm:"not null default 0 comment('是否删除，0删除 1：未删除') TINYINT(3)"`
}
