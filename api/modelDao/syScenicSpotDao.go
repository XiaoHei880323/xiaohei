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

type SyScenicSpotDao struct {
	engine *xorm.Engine
}

func NewSyScenicSpotDao(engine *xorm.Engine) SyScenicSpotDao {
	return SyScenicSpotDao{engine: engine}
}

func (SyScenicSpotDao) TableName() string { return "sy_scenic_spot" }

func (m SyScenicSpotDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
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

func (m SyScenicSpotDao) GetList(whereSlice []interface{}, page int, pageSize int) (int, []*model.SyScenicSpot, error) {
	var list []*model.SyScenicSpot
	params := whereSlice[1:]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy("id desc").Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

func (m SyScenicSpotDao) GetInfo(whereSlice []interface{}) (bool, *model.SyScenicSpot, error) {
	item := new(model.SyScenicSpot)
	params := whereSlice[1:]
	has, err := m.engine.Where(whereSlice[0], params...).Get(item)
	return has, item, err
}

func (m SyScenicSpotDao) Update(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
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
