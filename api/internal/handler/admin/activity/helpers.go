package activity

import (
	reponse "api/comment/response"
	"api/resp"
	"net/http"
	"strconv"
)

func parseErr(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
}

func uidFromHeader(r *http.Request) int {
	v, _ := strconv.Atoi(r.Header.Get("uid"))
	return v
}
