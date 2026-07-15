package reponse

import (
	"api/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(w http.ResponseWriter, resp *resp.CommonReply, r *http.Request, err error) {
	var body Body
	body.Code = int(resp.Code)
	body.Msg = resp.Message
	if nil != err {
		body.Msg += "," + err.Error()
	}
	body.Data = resp.Data

	httpx.OkJson(w, body)
}

func ReturnStruct() *resp.CommonReply {
	response := &resp.CommonReply{
		Code:    200,
		Message: "获取成功",
	}
	return response
}
