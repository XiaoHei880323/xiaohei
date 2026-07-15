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

type SyActivityConfigDao struct {
	engine *xorm.Engine
}

func NewSyActivityConfigDao(engine *xorm.Engine) SyActivityConfigDao {
	return SyActivityConfigDao{engine: engine}
}

func (SyActivityConfigDao) TableName() string { return "sy_activity_config" }

func (m SyActivityConfigDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
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

func (m SyActivityConfigDao) GetList(whereSlice []interface{}, page int, pageSize int) (int, []*model.SyActivityConfig, error) {
	var list []*model.SyActivityConfig
	params := whereSlice[1:]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy("id desc").Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

func (m SyActivityConfigDao) GetInfo(whereSlice []interface{}) (bool, *model.SyActivityConfig, error) {
	item := new(model.SyActivityConfig)
	params := whereSlice[1:]
	has, err := m.engine.Where(whereSlice[0], params...).Get(item)
	return has, item, err
}

func (m SyActivityConfigDao) Update(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
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

// UpdateAll 不检查影响行数（用于批量更新，可能影响0行）
func (m SyActivityConfigDao) UpdateAll(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(whereUpdate, updateMap, m.TableName())
	if err != nil {
		return errors.New(fmt.Sprintf("update sql err[%v]", err.Error()))
	}
	if session == nil {
		_, err = m.engine.Exec(res...)
	} else {
		_, err = session.Exec(res...)
	}
	return err
}
