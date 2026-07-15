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

type SyPointsRoleDao struct {
	engine *xorm.Engine
}

func NewSyPointsRoleDao(engine *xorm.Engine) SyPointsRoleDao {
	return SyPointsRoleDao{
		engine: engine,
	}
}

func (SyPointsRoleDao) TableName() string {
	return "sy_points_role"
}

type SyPointsRoleAndGood struct {
	model.SyPointsRole `xorm:"extends"`
	model.SyGood       `xorm:"extends"`
}

func (SyPointsRoleAndGood) TableName() string {
	return "sy_points_role"
}

func (m SyPointsRoleDao) GetPointsRoleAndGood(whereSlice []interface{}, page, pageSize int, orderBy string) (count int, list []*SyPointsRoleAndGood, err error) {
	params := whereSlice[1:len(whereSlice)]
	countInt64, err := m.engine.Alias("p").
		Join("left", "sy_good as g", "p.good_id = g.id").
		Where(whereSlice[0], params...).OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	count = int(countInt64)
	return
}

// 插入操作
func (m SyPointsRoleDao) Insert(params []map[string]interface{}, sessionParams *xorm.Session) error {
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
func (s SyPointsRoleDao) Update(whereUpdate []interface{}, paramsUpdate map[string]interface{}, session *xorm.Session) error {
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
func (m SyPointsRoleDao) GetInfo(whereSlice []interface{}) (bool, *model.SyPointsRole, error) {
	good := new(model.SyPointsRole)
	params := whereSlice[1:len(whereSlice)]
	has, err := m.engine.Where(whereSlice[0], params...).Get(good)
	return has, good, err
}
