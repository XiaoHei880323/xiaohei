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

type SySigninActivityScenicDao struct {
	engine *xorm.Engine
}

func NewSySigninActivityScenicDao(engine *xorm.Engine) SySigninActivityScenicDao {
	return SySigninActivityScenicDao{engine: engine}
}

func (SySigninActivityScenicDao) TableName() string { return "sy_signin_activity_scenic" }

func (m SySigninActivityScenicDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
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

func (m SySigninActivityScenicDao) GetList(whereSlice []interface{}, page int, pageSize int) (int, []*model.SySigninActivityScenic, error) {
	var list []*model.SySigninActivityScenic
	params := whereSlice[1:]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy("id asc").Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

func (m SySigninActivityScenicDao) GetInfo(whereSlice []interface{}) (bool, *model.SySigninActivityScenic, error) {
	item := new(model.SySigninActivityScenic)
	params := whereSlice[1:]
	has, err := m.engine.Where(whereSlice[0], params...).Get(item)
	return has, item, err
}

func (m SySigninActivityScenicDao) Update(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
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

// SumPointsByActivity 计算某活动下所有启用景点的积分之和
func (m SySigninActivityScenicDao) SumPointsByActivity(activityId int) (int, error) {
	type Result struct {
		Total int `xorm:"total"`
	}
	var r Result
	_, err := m.engine.SQL(
		"SELECT COALESCE(SUM(sign_points),0) AS total FROM sy_signin_activity_scenic WHERE activity_id = ? AND status = 1 AND is_delete = 0",
		activityId,
	).Get(&r)
	return r.Total, err
}
