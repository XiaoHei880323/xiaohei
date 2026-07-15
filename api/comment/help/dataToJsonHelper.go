package helper

import "encoding/json"

func ConvertStrutToJson(stl interface{}) string {
	str, err := json.Marshal(stl)
	if err != nil {
		return ""
	}
	return string(str)
}
