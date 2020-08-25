package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-api-boilerplate/pkg/ecode"
	"gin-api-boilerplate/pkg/log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//写入Response信息
func WriteResponse(w http.ResponseWriter, data interface{}, e error) {
	err := ecode.Cause(e)
	res := Response{
		Code:    err.Code(),
		Message: err.Message(),
		Data:    data,
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err1 := enc.Encode(res); err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

}

//发送response信息
func SendResponse(c *gin.Context, data interface{}, err error) {
	e := ecode.Cause(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    e.Code(),
		Message: e.Message(),
		Data:    data,
	})
}

//成功
func Success(c *gin.Context, data interface{}) {
	SendResponse(c, data, nil)
}

//错误
func Error(c *gin.Context, err error) {
	log.Sugar().Debugf("Error", err.Error())
	SendResponse(c, nil, err)
}

//带数据的错误
func ErrorWithData(c *gin.Context, err error, data interface{}) {
	SendResponse(c, data, err)
}
