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

type StudentDao struct {
	engine *xorm.Engine
}

func NewStudentDao(engine *xorm.Engine) StudentDao {
	return StudentDao{engine: engine}
}

func (StudentDao) TableName() string { return "course_student" }

func (m StudentDao) Insert(params []map[string]interface{}, session *xorm.Session) error {
	if len(params) == 0 {
		return fmt.Errorf("no data to insert [%v]", helper.InterfaceHelperObject.ToString(params))
	}
	res, err := dbs.InsertHelperObject.GetInsertSql(params, m.TableName())
	if err != nil {
		return fmt.Errorf("build student insert sql: %w", err)
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
		return errors.New("新增学生数量为0")
	}
	return nil
}

func (m StudentDao) Update(where []interface{}, values map[string]interface{}, session *xorm.Session) error {
	res, err := dbs.UpdateHelperObject.Update(where, values, m.TableName())
	if err != nil {
		return fmt.Errorf("build student update sql: %w", err)
	}
	if len(res) == 0 {
		return errors.New("没有可更新的学生数据")
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
		return errors.New("修改学生数量为0")
	}
	return nil
}

func (m StudentDao) GetList(where []interface{}, page, pageSize int, orderBy string) (int, []*model.Student, error) {
	list := make([]*model.Student, 0)
	params := where[1:]
	if orderBy == "" {
		orderBy = "id desc"
	}
	count, err := m.engine.Table(m.TableName()).Where(where[0], params...).OrderBy(orderBy).
		Limit(pageSize, pageSize*(page-1)).FindAndCount(&list)
	return int(count), list, err
}

func (m StudentDao) GetInfo(where []interface{}) (bool, *model.Student, error) {
	student := new(model.Student)
	params := where[1:]
	has, err := m.engine.Table(m.TableName()).Where(where[0], params...).Get(student)
	return has, student, err
}
