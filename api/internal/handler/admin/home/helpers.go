package home

import (
	reponse "api/comment/response"
	"api/resp"
	"net/http"
)

func parseErr(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
}
