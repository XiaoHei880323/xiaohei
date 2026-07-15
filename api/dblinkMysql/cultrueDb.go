package dblinkMysql

import (
	"api/modelDao"
	"xorm.io/xorm"
)

type CultrueDbStruct struct {
	SyAdminDao                modelDao.SyAdminDao
	SyActivityDao             modelDao.SyActivityDao
	SyGoodDao                 modelDao.SyGoodDao
	SyPointsRoleDao           modelDao.SyPointsRoleDao
	SyUserDao                 modelDao.SyUserDao
	SyUserPointsDao           modelDao.SyUserPointsDao
	SySeckillActivityGoodDao  modelDao.SySeckillActivityGoodDao
	SyScenicSpotDao           modelDao.SyScenicSpotDao
	SySigninActivityScenicDao modelDao.SySigninActivityScenicDao
	SyHomeConfigDao           modelDao.SyHomeConfigDao
	SyNoticeDao               modelDao.SyNoticeDao
	SyActivityConfigDao       modelDao.SyActivityConfigDao
	SyActivityConfigItemDao   modelDao.SyActivityConfigItemDao
	StudentDao                modelDao.StudentDao
	TeacherDao                modelDao.TeacherDao
	CourseMainDao             modelDao.CourseMainDao
	CourseEvaluationDao       modelDao.CourseEvaluationDao
	CourseMediaDao            modelDao.CourseMediaDao
	CourseErrorCollectionDao  modelDao.CourseErrorCollectionDao
}

func NewCultrueDbStruct(engine *xorm.Engine) *CultrueDbStruct {
	return &CultrueDbStruct{
		SyAdminDao:                modelDao.NewSyAdminDao(engine),
		SyActivityDao:             modelDao.NewSyActivityDao(engine),
		SyGoodDao:                 modelDao.NewSyGoodDao(engine),
		SyPointsRoleDao:           modelDao.NewSyPointsRoleDao(engine),
		SyUserDao:                 modelDao.NewSyUserDao(engine),
		SyUserPointsDao:           modelDao.NewSyUserPointsDao(engine),
		SySeckillActivityGoodDao:  modelDao.NewSySeckillActivityGoodDao(engine),
		SyScenicSpotDao:           modelDao.NewSyScenicSpotDao(engine),
		SySigninActivityScenicDao: modelDao.NewSySigninActivityScenicDao(engine),
		SyHomeConfigDao:           modelDao.NewSyHomeConfigDao(engine),
		SyNoticeDao:               modelDao.NewSyNoticeDao(engine),
		SyActivityConfigDao:       modelDao.NewSyActivityConfigDao(engine),
		SyActivityConfigItemDao:   modelDao.NewSyActivityConfigItemDao(engine),
		StudentDao:                modelDao.NewStudentDao(engine),
		TeacherDao:                modelDao.NewTeacherDao(engine),
		CourseMainDao:             modelDao.NewCourseMainDao(engine),
		CourseEvaluationDao:       modelDao.NewCourseEvaluationDao(engine),
		CourseMediaDao:            modelDao.NewCourseMediaDao(engine),
		CourseErrorCollectionDao:  modelDao.NewCourseErrorCollectionDao(engine),
	}
}
