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

type CourseMainDao struct {
	engine *xorm.Engine
}

func NewCourseMainDao(engine *xorm.Engine) CourseMainDao {
	return CourseMainDao{engine: engine}
}

func (CourseMainDao) TableName() string { return "course_main" }

func (m CourseMainDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_main insert sql: %w", err)
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
		return errors.New("新增课程数量为0")
	}
	return nil
}

func (m CourseMainDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_main update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的课程数据")
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
		return errors.New("修改课程数量为0")
	}
	return nil
}

func (m CourseMainDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.CourseMain, error) {
	list := make([]*model.CourseMain, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m CourseMainDao) GetInfo(where []interface{}) (bool, *model.CourseMain, error) {
	item := new(model.CourseMain)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(item)
	return has, item, err
}
