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

type TeacherDao struct {
	engine *xorm.Engine
}

func NewTeacherDao(engine *xorm.Engine) TeacherDao {
	return TeacherDao{engine: engine}
}

func (TeacherDao) TableName() string { return "course_teacher" }

func (m TeacherDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build teacher insert sql: %w", err)
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
		return errors.New("新增教师数量为0")
	}
	return nil
}

func (m TeacherDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build teacher update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的教师数据")
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
		return errors.New("修改教师数量为0")
	}
	return nil
}

func (m TeacherDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.Teacher, error) {
	list := make([]*model.Teacher, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m TeacherDao) GetInfo(where []interface{}) (bool, *model.Teacher, error) {
	teacher := new(model.Teacher)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(teacher)
	return has, teacher, err
}
