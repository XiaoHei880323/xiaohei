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

type CourseMediaDao struct {
	engine *xorm.Engine
}

func NewCourseMediaDao(engine *xorm.Engine) CourseMediaDao {
	return CourseMediaDao{engine: engine}
}

func (CourseMediaDao) TableName() string { return "course_media" }

func (m CourseMediaDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_media insert sql: %w", err)
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
		return errors.New("新增媒体资源数量为0")
	}
	return nil
}

func (m CourseMediaDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build course_media update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的媒体资源数据")
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
		return errors.New("修改媒体资源数量为0")
	}
	return nil
}

func (m CourseMediaDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.CourseMedia, error) {
	list := make([]*model.CourseMedia, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m CourseMediaDao) GetInfo(where []interface{}) (bool, *model.CourseMedia, error) {
	item := new(model.CourseMedia)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(item)
	return has, item, err
}
