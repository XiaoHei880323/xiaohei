package modelDao

import (
	helper "api/comment/help"
	"api/comment/help/dbs"
	"api/model"
	"database/sql"
	"errors"
	"fmt"
	"xorm.io/xorm"
)

type SyAdminDao struct {
	engine *xorm.Engine
}

func NewSyAdminDao(engine *xorm.Engine) SyAdminDao {
	return SyAdminDao{
		engine: engine,
	}
}

func (SyAdminDao) TableName() string {
	return "sy_admin"
}

// 插入操作
func (m SyAdminDao) Insert(params []map[string]interface{}, sessionParams *xorm.Session) error {
	if len(params) == 0 {
		return errors.New(fmt.Sprintf("No data to Insert [%v]", helper.InterfaceHelperObject.ToString(params)))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if nil != err {
		return errors.New(fmt.Sprintf("Failed to perform insert...err[%v] res[%v]", err.Error(), helper.InterfaceHelperObject.ToString(res)))
	}
	var tmp sql.Result
	var errTmp error
	if nil == sessionParams {
		tmp, errTmp = m.engine.Exec(res...)
	} else {
		tmp, errTmp = sessionParams.Exec(res...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {

		return errors.New("添加数量为0")
	}
	return nil

}

// 修改
func (s SyAdminDao) Update(whereUpdate []interface{}, paramsUpdate map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(whereUpdate, paramsUpdate, s.TableName())
	if nil != err {
		return errors.New(fmt.Sprintf("Failed to stitch basic data...err[%v] res[%v]", err.Error(), helper.InterfaceHelperObject.ToString(res)))
	}
	if len(res) == 0 {
		return errors.New(fmt.Sprintf("No data to update [%v]", helper.InterfaceHelperObject.ToString(res)))
	}
	var tmp sql.Result
	var errTmp error
	if nil == session {
		tmp, errTmp = s.engine.Exec(res...)
	} else {
		tmp, errTmp = session.Exec(res...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {

		return errors.New("修改数量为0")
	}
	return nil
}

// 查询详情
func (m SyAdminDao) GetInfo(whereSlice []interface{}) (bool, *model.SyAdmin, error) {
	adminUser := new(model.SyAdmin)
	params := whereSlice[1:len(whereSlice)]
	has, err := m.engine.Where(whereSlice[0], params...).Get(adminUser)
	return has, adminUser, err
}

// 查询分页的数据
func (m SyAdminDao) GetCorList(whereSlice []interface{}, page int, pageSize int, orderBy string) (count int, list []*model.SyAdmin, err error) {
	params := whereSlice[1:len(whereSlice)]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

// 查询数量
func (m SyAdminDao) GetCount(whereSlice []interface{}) (count int, err error) {
	list := new(model.SyAdmin)
	params := whereSlice[1:len(whereSlice)]
	countInt64, err := m.engine.Where(whereSlice[0], params...).Count(list)
	return int(countInt64), err
}

// 获取用户的数据信息in
func (m SyAdminDao) GetAdminWhereInInt(key string, values []int) (list []*model.SyAdmin, err error) {
	err = m.engine.In(key, values).Find(&list)
	return list, err
}
