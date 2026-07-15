package dbs

import (
	helper "api/comment/help"
	"errors"

	"fmt"
	"strings"
)

/**
*  UpdateHelper
*  @Description: Update封装
 */
type InsertHelper struct {
}

var InsertHelperObject InsertHelper

/**
 *  Insert
 *  @Description:
 *  @receiver receiver
 *  @param insertFiledMap
 *  @param tableName
 *  @return []interface{}
 *  @return error
 */
//func (receiver *InsertHelper) Insert(insertFiledMap []map[string]interface{}, tableName string) ([]interface{}, error) {
//	returnInterface := make([]interface{}, 0)
//	if len(insertFiledMap) == 0 {
//		return returnInterface, errors.New(fmt.Sprintf("The data to be inserted is empty... insertFiledMap[%v]", helper.InterfaceHelperObject.ToString(insertFiledMap)))
//	}
//	//更新条件必传防止全盘更新误操作
//	if len(insertFiledMap[0]) == 0 {
//		return returnInterface, errors.New(fmt.Sprintf("The field to be inserted is empty... insertFiledMap[%v]", helper.InterfaceHelperObject.ToString(insertFiledMap)))
//	}
//	sql := fmt.Sprintf("INSERT INTO %v (", tableName)
//	//拼接字符串
//	keySlice := make([]string, 0)
//	for filed, _ := range insertFiledMap[0] {
//		sql = fmt.Sprintf("%v%v?,", sql, filed)
//		keySlice = append(keySlice, filed)
//	}
//	//删除最后一个字符
//	sql = sql[:len(sql)-1] + ") VALUES"
//	//问号对应的数据
//	InsertFiled := make([]interface{}, 0)
//	//拼接子
//	itemStr := "("
//	for _, item := range keySlice {
//		itemStr += "?,"
//		fmt.Sprintf(item)
//	}
//	itemStr += ")"
//
//	allSlice := make([]string, 0)
//	for _, item := range insertFiledMap {
//		allSlice = append(allSlice, itemStr)
//		for _, i2 := range keySlice {
//			InsertFiled = append(InsertFiled, item[i2])
//		}
//	}
//	sql += strings.Join(allSlice, ",")
//	//处理数据
//	returnInterface = append(returnInterface, sql)
//	returnInterface = append(returnInterface, InsertFiled...)
//	return returnInterface, nil
//}

func (h *InsertHelper) GetInsertSql(paramsSlice []map[string]interface{}, tableName string) ([]interface{}, error) {
	res := make([]interface{}, 0)
	getStr := func(data interface{}) string {
		return helper.InterfaceHelperObject.ToString(data)
	}
	if 0 == len(paramsSlice) {
		return res, errors.New(fmt.Sprintf("待插入的数据不允许为空;%s", getStr(paramsSlice)))
	}
	if "" == tableName {
		return res, errors.New(fmt.Sprintf("待插入的表名不允许为空;%s", getStr(paramsSlice)))
	}
	keySlice := make([]string, 0)
	for key, _ := range paramsSlice[0] {
		keySlice = append(keySlice, key)
	}
	if len(keySlice) < 1 {
		return res, errors.New(fmt.Sprintf("待插入的数据key不允许为空;%s", getStr(paramsSlice)))
	}
	//sqlBase := "INSERT INTO "+tableName+" (`project_id`,`url`,`type`) VALUES "
	sqlBase := "INSERT INTO " + tableName + " (`" + strings.Join(keySlice, "`,`") + "`) VALUES "
	params := make([]interface{}, 0)
	tmp1 := make([]string, 0)
	for _, i2 := range paramsSlice {
		tmpItem := make([]string, 0)
		for _, i4 := range keySlice {
			value, has := i2[i4]
			if !has {
				return res, errors.New(fmt.Sprintf("待插入的数据key无对应的值;%s", getStr(paramsSlice)))
			}
			tmpItem = append(tmpItem, "?")
			params = append(params, value)
		}
		tmpString := "(" + strings.Join(tmpItem, ",") + ")"
		tmp1 = append(tmp1, tmpString)
	}
	sqlBase += strings.Join(tmp1, ",")
	res = append(res, sqlBase)
	res = append(res, params...)
	return res, nil
}

//插入
//func (h *InsertHelper) SQlInsert(paramsSlice []map[string]interface{}, tableName string) (inserId int , count int, error) {
//	res , err := InsertHelperObject.GetInsertSql(paramsSlice , tableName)
//	if err != nil{
//		return res , err
//	}
//}
