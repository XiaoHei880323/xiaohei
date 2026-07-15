package modelDao

import (
	helper "api/comment/help"
	"api/comment/help/dbs"
	"database/sql"
	"errors"
	"fmt"

	//helper "api/comment/help"
	//"api/comment/help/dbs"
	"api/model"
	//"database/sql"
	"xorm.io/xorm"
)

type SyActivityDao struct {
	engine *xorm.Engine
}

func NewSyActivityDao(engine *xorm.Engine) SyActivityDao {
	return SyActivityDao{
		engine: engine,
	}
}

func (SyActivityDao) TableName() string {
	return "sy_activity"
}

func (m SyActivityDao) GetCorList(whereSlice []interface{}, page int, pageSize int, orderBy string) (count int, list []*model.SyActivity, err error) {
	params := whereSlice[1:]
	findSql := m.engine.Where(whereSlice[0], params...)
	findSql = findSql.Where("is_delete = ?", 1)
	countInt64, err := findSql.OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

// 新增活动
func (m SyActivityDao) Insert(params []map[string]interface{}, sessionParams *xorm.Session) error {
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

// 查询详情
func (m SyActivityDao) GetInfo(whereSlice []interface{}) (bool, *model.SyActivity, error) {
	adminUser := new(model.SyActivity)
	params := whereSlice[1:len(whereSlice)]
	has, err := m.engine.Where(whereSlice[0], params...).Get(adminUser)
	return has, adminUser, err
}

// 根据ID列表批量查询活动
func (m SyActivityDao) GetWhereInList(key string, values []int) ([]*model.SyActivity, error) {
	var list []*model.SyActivity
	err := m.engine.In(key, values).Find(&list)
	return list, err
}

// 修改
func (s SyActivityDao) Update(whereUpdate []interface{}, paramsUpdate map[string]interface{}, session *xorm.Session) error {
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
