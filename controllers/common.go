package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data  interface{} `json:"data"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}, data interface{}) {
	//json := &JsonStruct{Code: code, Msg: msg, Data: data}
	json := &JsonErrStruct{Code: code, Msg: msg, Data: data}
	c.JSON(http.StatusInternalServerError, json)
}
