package controllers

import (
	"fmt"
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 让每一个controller下面能定义同名函数，使用结构体
type UserController struct{}

// 方法调用：方法是与结构体（或自定义类型）相关联的函数，它们可以通过结构体的实例进行调用。
func (u UserController) GetUserInfo(c *gin.Context) {
	// get id
	idStr := c.Param("id")
	//name := c.Param("name")
	fmt.Println(">>>>", idStr)  // for debug

	id, _ := strconv.Atoi(idStr)
    // fixme, report error
	user,_ := models.GetUserTest(id)

	fmt.Println(user.Name)
	fmt.Println(user.Id)

	// 因为在同一个包中，所以可以直接调用
	ReturnSuccess(c, http.StatusOK, idStr, user, 1)
	//ReturnSuccess(c, http.StatusOK, "request success", id, 1)
	//ReturnSuccess(c, http.StatusOK, "request success", nil, 1)   // "data":null
}

func (u UserController) GetList(c *gin.Context) {
	//logger.Write("日志信息","user")
	// catch exception
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()

	// test case exception
	num1 := 1
	num2 := 0
	num3 := num1 / num2

	fmt.Println("[D] result: ", num3)

	ReturnError(c, 40000, "no data", nil)
}
