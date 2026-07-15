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

type CourseEvaluationDao struct {
	engine *xorm.Engine
}

func NewCourseEvaluationDao(engine *xorm.Engine) CourseEvaluationDao {
	return CourseEvaluationDao{engine: engine}
}

func (CourseEvaluationDao) TableName() string { return "course_evaluation" }

func (m CourseEvaluationDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_evaluation insert sql: %w", err)
	}
	var result sql.Result
	if session == nil {
		result, err = m.engine.Exec(res...)
	} else {
		result, err = session.Exec(res...)
	}
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("新增评价数量为0")
	}
	return nil
}

func (m CourseEvaluationDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_evaluation update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的评价数据")
	}
	var result sql.Result
	if session == nil {
		result, err = m.engine.Exec(res...)
	} else {
		result, err = session.Exec(res...)
	}
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("修改评价数量为0")
	}
	return nil
}

func (m CourseEvaluationDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.CourseEvaluation, error) {
	list := make([]*model.CourseEvaluation, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m CourseEvaluationDao) GetInfo(where []interface{}) (bool, *model.CourseEvaluation, error) {
	item := new(model.CourseEvaluation)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(item)
	return has, item, err
}
