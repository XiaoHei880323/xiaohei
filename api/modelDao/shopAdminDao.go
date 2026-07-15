package modelDao

import (
	"api/model"
	"xorm.io/xorm"
)

type ShopAdminDao struct {
	engine *xorm.Engine
}

func NewShopAdminDao(engine *xorm.Engine) ShopAdminDao {
	return ShopAdminDao{
		engine: engine,
	}
}

func (ShopAdminDao) TableName() string {
	return "shop_admin"
}

func (m ShopAdminDao) GetCorList(whereSlice []interface{}, page int, pageSize int, orderBy string) (count int, list []*model.ShopAdmin, err error) {
	params := whereSlice[1:len(whereSlice)]
	if orderBy == "" {
		orderBy = "id desc"
	}
	countInt64, err := m.engine.Where(whereSlice[0], params...).
		OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).
		FindAndCount(&list)
	return int(countInt64), list, err
}
