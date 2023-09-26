package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct{}

func (o OrderController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, http.StatusOK, "request success order", gin.H{}, 1)
}

func (o OrderController) GetList(c *gin.Context) {
	ReturnError(c, 40000, "no data", nil)
}
