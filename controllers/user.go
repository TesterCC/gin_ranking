package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 让每一个controller下面能定义同名函数，使用结构体
type UserController struct{}

// 方法调用：方法是与结构体（或自定义类型）相关联的函数，它们可以通过结构体的实例进行调用。
func (u UserController) GetUserInfo(c *gin.Context) {
	// 因为在同一个包中，所以可以直接调用
	ReturnSuccess(c, http.StatusOK, "request success", gin.H{}, 1)
	//ReturnSuccess(c, http.StatusOK, "request success", nil, 1)   // "data":null
}

func (u UserController) GetList(c *gin.Context) {
	ReturnError(c, 40000, "no data", nil)
}
