package resData

import (
	"TestGOWeb/recordlog"
	"encoding/json"
	"net/http"
)

var ErrCodeMap map[string]int32

var ErrDataErr string
var ErrPasswdErr string
var ErrParamentErr string

func init() {
	init_ErrType()
	init_ErrCode()

}

type RespData struct {
	Code      int32       `json:"Code"`
	IsSuccess bool        `json:"IsSuccess"`
	Message   string      `json:"Msg"`
	Data      interface{} `json:"Data"`
}

func Write(w http.ResponseWriter, err string, data interface{}) {
	code, ok := ErrCodeMap[err]
	recordlog.Debug("code：", code)
	recordlog.Debug("err: ", err)
	rsp := RespData{
		Code:    code,
		Data:    data,
		Message: err,
	}
	if !ok {
		rsp.IsSuccess = true
	} else {
		rsp.IsSuccess = false
	}
	writeContext, _ := json.Marshal(rsp)
	w.Write(writeContext)
}

func init_ErrType() {
	ErrDataErr = "数据异常" //数据异常
	ErrParamentErr = "参数异常"
	ErrPasswdErr = "密码错误"
}

func init_ErrCode() {
	ErrCodeMap = make(map[string]int32)
	ErrCodeMap[ErrParamentErr] = 501
	ErrCodeMap[ErrDataErr] = 502
	ErrCodeMap[ErrPasswdErr] = 503
}
