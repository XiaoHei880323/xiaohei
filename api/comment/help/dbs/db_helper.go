package dbs

/**
*  DbHelper
*  @Description:
 */
type DbHelper struct {
}

/**
*  DbHelper
*  @Description:
 */
var DbHelperObject DbHelper

/**
 *  SliceToString
 *  @Description: string切面转 where 处理 mysql 中 in的场景
 *  @receiver m
 *  @param slice
 *  @return string
 */
//func (m *DbHelper) SliceToString(slice *[]string) string {
//
//	sliceTran := make([]string, 0)
//	sliceMap := make(map[string]string)
//	result := ""
//	if 0 == len(*slice) {
//		return result
//	}
//	//大小判断并去重
//	for k, v := range *slice {
//		sliceMap[v] = string(k)
//	}
//	for k1, _ := range sliceMap {
//		sliceTran = append(sliceTran, string(k1))
//	}
//	tranLength := len(sliceTran)
//	if 0 == len(sliceTran) {
//		return result
//	}
//	//拼接查询字符串
//	in := ""
//	for k3, v3 := range sliceTran {
//		if v3 == "" {
//			v3 = "''"
//		}
//		if 0 == k3 {
//			in = fmt.Sprintf("%v%v%v", "(", in, v3)
//		} else if tranLength-1 == k3 {
//			in = fmt.Sprintf("%v%v%v%v", in, ",", v3, ")")
//		} else {
//			in = fmt.Sprintf("%v%v%v", in, ",", v3)
//		}
//	}
//	where := fmt.Sprintf("%v", in)
//	return where
//}
