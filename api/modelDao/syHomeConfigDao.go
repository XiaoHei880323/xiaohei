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

type SyHomeConfigDao struct {
	engine *xorm.Engine
}

func NewSyHomeConfigDao(engine *xorm.Engine) SyHomeConfigDao {
	return SyHomeConfigDao{engine: engine}
}

func (SyHomeConfigDao) TableName() string { return "sy_home_config" }

func (m SyHomeConfigDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return errors.New(fmt.Sprintf("No data to Insert [%v]", helper.InterfaceHelperObject.ToString(params)))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return errors.New(fmt.Sprintf("insert sql err[%v]", err.Error()))
	}
	var tmp sql.Result
	var errTmp error
	if session == nil {
		tmp, errTmp = m.engine.Exec(res...)
	} else {
		tmp, errTmp = session.Exec(res...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, _ := tmp.RowsAffected()
	if count < 1 {
		return errors.New("添加数量为0")
	}
	return nil
}

func (m SyHomeConfigDao) GetList(whereSlice []interface{}, page int, pageSize int) (int, []*model.SyHomeConfig, error) {
	var list []*model.SyHomeConfig
	params := whereSlice[1:]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy("sort asc, id desc").Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

func (m SyHomeConfigDao) GetInfo(whereSlice []interface{}) (bool, *model.SyHomeConfig, error) {
	item := new(model.SyHomeConfig)
	params := whereSlice[1:]
	has, err := m.engine.Where(whereSlice[0], params...).Get(item)
	return has, item, err
}

func (m SyHomeConfigDao) Update(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(whereUpdate, updateMap, m.TableName())
	if err != nil {
		return errors.New(fmt.Sprintf("update sql err[%v]", err.Error()))
	}
	var tmp sql.Result
	var errTmp error
	if session == nil {
		tmp, errTmp = m.engine.Exec(res...)
	} else {
		tmp, errTmp = session.Exec(res...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, _ := tmp.RowsAffected()
	if count < 1 {
		return errors.New("修改数量为0")
	}
	return nil
}
