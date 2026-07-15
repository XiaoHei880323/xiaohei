package dblinkMysql

import (
	"api/modelDao"
	"xorm.io/xorm"
)

type ShopDbStruct struct {
	ShopAdminDao modelDao.ShopAdminDao
}

func NewShopDbStruct(engine *xorm.Engine) *ShopDbStruct {
	return &ShopDbStruct{
		ShopAdminDao: modelDao.NewShopAdminDao(engine),
	}
}
