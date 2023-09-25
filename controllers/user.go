package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context)  {
	// 因为在同一个包中，所以可以直接调用
	ReturnSuccess(c, http.StatusOK, "request success", nil, 1)
}