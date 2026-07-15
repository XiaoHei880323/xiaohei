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

type SyUserPointsDao struct {
	engine *xorm.Engine
}

func NewSyUserPointsDao(engine *xorm.Engine) SyUserPointsDao {
	return SyUserPointsDao{
		engine: engine,
	}
}
func (SyUserPointsDao) Table() string {
	return "sy_user_points"
}

// 插入操作
func (m SyUserPointsDao) Insert(params []map[string]interface{}, sessionParams *xorm.Session) error {
	if len(params) == 0 {
		return errors.New(fmt.Sprintf("No data to Insert [%v]", helper.InterfaceHelperObject.ToString(params)))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.Table())
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
func (s SyUserPointsDao) Update(whereUpdate []interface{}, paramsUpdate map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(whereUpdate, paramsUpdate, s.Table())
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

type SumPointsDao struct {
	Points int `json:"points"`
}

func (SumPointsDao) Table() string {
	return "sy_user_points"
}

// 查询用户的积分
func (m SyUserPointsDao) GetUserPointsSum(where []interface{}) (sumPoints int, err error) {
	sumPointsDao := new(SumPointsDao)
	params := where[1:len(where)]
	sumInt64, err := m.engine.Table("sy_user_points").Where(where[0], params...).SumInt(sumPointsDao, "points")
	sumPoints = int(sumInt64)
	return
}

// 查询分页的数据
func (m SyUserPointsDao) GetUserPointsListPage(whereSlice []interface{}, page int, pageSize int, orderBy string) (count int, list []*model.SyUserPoints, err error) {
	params := whereSlice[1:len(whereSlice)]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy(orderBy).Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}
