package dbs

import (
	"errors"
	"fmt"
	"strings"
)

/**
*  UpdateHelper
*  @Description: Update封装
 */
type UpdateHelper struct {
}

/**
*  UpdateHelperObject
*  @Description: Update封装
 */
var UpdateHelperObject UpdateHelper

/**
 *  updateBatchByIds
 *  @Description:
 *  @receiver receiver
 *  @param paramsSet
 *  @param tableName
 *  @return string
 *  @return error
 */
func (receiver *UpdateHelper) UpdateBatchByIds(paramsSet []map[string]string, tableName string) ([]interface{}, error) {
	tmp := make([]interface{}, 0)
	resultTmp := make([]interface{}, 0)
	if 0 == len(paramsSet) {
		return resultTmp, nil
	}
	if 1 >= len(paramsSet[0]) {
		return resultTmp, errors.New("更新元素必须大于等于2个\n")
	}
	allSql := "update " + tableName + " SET "
	//循环拼接sql
	index := 0
	for key1, _ := range paramsSet[0] {
		if "id" == key1 {
			continue
		}
		itemString := ""
		if 0 != index {
			itemString = ","
		}
		itemString = fmt.Sprintf("%v`%v` = CASE ", itemString, key1)
		for _, value2 := range paramsSet {
			//判断key是否存在
			if val, ok := value2["id"]; !ok {
				return resultTmp, errors.New("key id 不存在\n")
				if val == "0" || val == "" {
					return resultTmp, errors.New("id 对应的value不能为 0 或者 空\n")
				}
			}
			if _, ok := value2[key1]; !ok {
				return resultTmp, errors.New(key1 + " 不存在\n")
			}
			//itemString = fmt.Sprintf("%v WHEN `id` = %v THEN '%v' ", itemString, value2["id"], value2[key1])
			itemString = fmt.Sprintf("%v WHEN `id` = %v THEN ? ", itemString, value2["id"])
			tmp = append(tmp, value2[key1])
		}
		itemString = itemString + " END"
		allSql = allSql + itemString
		index++
	}
	whereParams := make([]string, 0)
	for _, value := range paramsSet {
		whereParams = append(whereParams, "'"+value["id"]+"'")
	}
	//拼接wehere条件
	allSql = fmt.Sprintf("%v where `id` in (%v)", allSql, strings.Join(whereParams, ","))
	//拼接数据
	resultTmp = append(resultTmp, allSql)
	//添加数据
	for _, v := range tmp {
		resultTmp = append(resultTmp, v)
	}
	return resultTmp, nil
}

/**
 *  Update
 *  @Description: 拼接更新sql
 *  @receiver receiver
 *  @param where
 *  @param mapParams
 *  @param tableName
 *  @return string
 */
func (receiver *UpdateHelper) Update(whereUpdate []interface{}, updateFiledMap map[string]interface{}, tableName string) ([]interface{}, error) {
	returnInterface := make([]interface{}, 0)
	if len(updateFiledMap) == 0 {
		return returnInterface, errors.New("需要更新的数据为空")
	}
	//更新条件必传防止全盘更新误操作
	if len(whereUpdate) == 0 {
		return returnInterface, errors.New("需要更新的条件不能为空")
	}
	sql := fmt.Sprintf("UPDATE %v SET ", tableName)
	//拼接字符串
	updateFiled := make([]interface{}, 0)
	for filed, filedValue := range updateFiledMap {
		sql = fmt.Sprintf("%v %v = ?,", sql, filed)
		updateFiled = append(updateFiled, filedValue)
	}
	//删除最后一个字符
	sql = sql[:len(sql)-1]
	//拼接where
	whereSql := whereUpdate[0]
	whereParams := whereUpdate[1:]
	sql = fmt.Sprintf("%v WHERE %v", sql, whereSql)
	//处理数据
	returnInterface = append(returnInterface, sql)
	returnInterface = append(returnInterface, updateFiled...)
	returnInterface = append(returnInterface, whereParams...)
	return returnInterface, nil
}
