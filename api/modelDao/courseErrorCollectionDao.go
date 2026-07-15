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

type CourseErrorCollectionDao struct {
	engine *xorm.Engine
}

func NewCourseErrorCollectionDao(engine *xorm.Engine) CourseErrorCollectionDao {
	return CourseErrorCollectionDao{engine: engine}
}

func (CourseErrorCollectionDao) TableName() string { return "course_error_collection" }

func (m CourseErrorCollectionDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_error_collection insert sql: %w", err)
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
		return errors.New("新增错题数量为0")
	}
	return nil
}

func (m CourseErrorCollectionDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_error_collection update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的错题数据")
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
		return errors.New("修改错题数量为0")
	}
	return nil
}

func (m CourseErrorCollectionDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.CourseErrorCollection, error) {
	list := make([]*model.CourseErrorCollection, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m CourseErrorCollectionDao) GetInfo(where []interface{}) (bool, *model.CourseErrorCollection, error) {
	item := new(model.CourseErrorCollection)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(item)
	return has, item, err
}
