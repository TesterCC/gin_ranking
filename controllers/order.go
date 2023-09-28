package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid int `json:"cid"`
}

func (o OrderController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, http.StatusOK, "request success order", gin.H{}, 1)
}

func (o OrderController) GetList(c *gin.Context) {
	// 参数传递2 不需要修改路由   post_form
	cid := c.PostForm("cid")
	name := c.DefaultPostForm("name", "Tester")    // name字段未传时，默认是Tester
	//ReturnError(c, 40000, "no data", nil)
	ReturnSuccess(c, http.StatusOK, "request success order", cid + "-" +name, 1)
}


func (o OrderController) GetListAll(c *gin.Context) {
	// 参数传递3 不需要修改路由   json_data 需要定义结构体
	param := make(map[string]interface{})
	err := c.BindJSON(&param)

	if err == nil {
		ReturnSuccess(c, http.StatusOK, "request success order", param, 1)
		return
	}

	ReturnError(c, 40001, "post data err", gin.H{"err": err})

}


func (o OrderController) GetListAllSearch(c *gin.Context) {
	// 参数传递4 定义结构体接收参数 不需要修改路由  json_data 需要定义结构体
	search := &Search{}
	err := c.BindJSON(&search)

	data := gin.H{
		"id": search.Cid,
		"name": search.Name,
	}

	if err == nil {
		ReturnSuccess(c, http.StatusOK, "request success order", data, 1)
		return
	}

	ReturnError(c, 40001, "post data err", gin.H{"err": err})

}