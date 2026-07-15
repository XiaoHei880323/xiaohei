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

type SyGoodDao struct {
	engine *xorm.Engine
}

func NewSyGoodDao(engine *xorm.Engine) SyGoodDao {
	return SyGoodDao{
		engine: engine,
	}
}

func (SyGoodDao) TableName() string {
	return "sy_good"
}

// 插入操作
func (m SyGoodDao) Insert(params []map[string]interface{}, sessionParams *xorm.Session) error {
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

// 获取列表
func (m SyGoodDao) GetGoodList(whereSlice []interface{}, page int, pageSize int, orderBy string) (count int, list []*model.SyGood, err error) {
	params := whereSlice[1:len(whereSlice)]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

// 查询详情
func (m SyGoodDao) GetInfo(whereSlice []interface{}) (bool, *model.SyGood, error) {
	good := new(model.SyGood)
	params := whereSlice[1:len(whereSlice)]
	has, err := m.engine.Where(whereSlice[0], params...).Get(good)
	return has, good, err
}

// 修改
func (s SyGoodDao) Update(whereUpdate []interface{}, paramsUpdate map[string]interface{}, session *xorm.Session) error {
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
