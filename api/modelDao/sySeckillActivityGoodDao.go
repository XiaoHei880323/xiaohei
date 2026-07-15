package modelDao

import (
	helper "api/comment/help"
	"api/comment/help/dbs"
	"api/model"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"xorm.io/xorm"
)

type SySeckillActivityGoodDao struct {
	engine *xorm.Engine
}

func NewSySeckillActivityGoodDao(engine *xorm.Engine) SySeckillActivityGoodDao {
	return SySeckillActivityGoodDao{engine: engine}
}

func (SySeckillActivityGoodDao) TableName() string {
	return "sy_seckill_activity_good"
}

// Insert 新增
func (m SySeckillActivityGoodDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return errors.New(fmt.Sprintf("No data to Insert [%v]", helper.InterfaceHelperObject.ToString(params)))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to perform insert...err[%v] res[%v]", err.Error(), helper.InterfaceHelperObject.ToString(res)))
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
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {
		return errors.New("添加数量为0")
	}
	return nil
}

// GetListByWhere 列表（where 条件切片，与其他 DAO 保持一致）
func (m SySeckillActivityGoodDao) GetListByWhere(whereSlice []interface{}, page int, pageSize int) (int, []*model.SySeckillActivityGood, error) {
	var list []*model.SySeckillActivityGood
	params := whereSlice[1:]
	countInt64, err := m.engine.Where(whereSlice[0], params...).OrderBy("id desc").Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(countInt64), list, err
}

// GetInfo 详情
func (m SySeckillActivityGoodDao) GetInfo(whereSlice []interface{}) (bool, *model.SySeckillActivityGood, error) {
	item := new(model.SySeckillActivityGood)
	params := whereSlice[1:]
	has, err := m.engine.Where(whereSlice[0], params...).Get(item)
	return has, item, err
}

// Update 修改
func (m SySeckillActivityGoodDao) Update(whereUpdate []interface{}, updateMap map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(whereUpdate, updateMap, m.TableName())
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to stitch basic data...err[%v] res[%v]", err.Error(), helper.InterfaceHelperObject.ToString(res)))
	}
	if len(res) == 0 {
		return errors.New(fmt.Sprintf("No data to update [%v]", helper.InterfaceHelperObject.ToString(res)))
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
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {
		return errors.New("修改数量为0")
	}
	return nil
}

// BatchUpdatePrice 批量修改秒杀价
func (m SySeckillActivityGoodDao) BatchUpdatePrice(ids []int, seckillPrice string, updateUid int, session *xorm.Session) error {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	placeholders := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	sqlStr := fmt.Sprintf("UPDATE %s SET seckill_price = ?, update_uid = ? WHERE id IN (%s) AND is_delete = 0",
		m.TableName(), placeholders)
	args := []interface{}{seckillPrice, updateUid}
	for _, id := range ids {
		args = append(args, id)
	}
	var tmp sql.Result
	var errTmp error
	if session == nil {
		tmp, errTmp = m.engine.Exec(append([]interface{}{sqlStr}, args...)...)
	} else {
		tmp, errTmp = session.Exec(append([]interface{}{sqlStr}, args...)...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {
		return errors.New("批量修改数量为0")
	}
	return nil
}

// BatchDelete 批量软删除
func (m SySeckillActivityGoodDao) BatchDelete(ids []int, updateUid int, session *xorm.Session) error {
	if len(ids) == 0 {
		return errors.New("ids不能为空")
	}
	placeholders := strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
	sqlStr := fmt.Sprintf("UPDATE %s SET is_delete = 1, update_uid = ? WHERE id IN (%s) AND is_delete = 0",
		m.TableName(), placeholders)
	args := []interface{}{updateUid}
	for _, id := range ids {
		args = append(args, id)
	}
	var tmp sql.Result
	var errTmp error
	if session == nil {
		tmp, errTmp = m.engine.Exec(append([]interface{}{sqlStr}, args...)...)
	} else {
		tmp, errTmp = session.Exec(append([]interface{}{sqlStr}, args...)...)
	}
	if errTmp != nil {
		return errTmp
	}
	count, errCount := tmp.RowsAffected()
	if errCount != nil {
		return errCount
	}
	if count < 1 {
		return errors.New("批量删除数量为0")
	}
	return nil
}
