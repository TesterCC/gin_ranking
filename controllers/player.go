package controllers

import (
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerController struct {}

func (p PlayerController) GetPlayers(c *gin.Context)  {
	// http://localhost:9999/player/list post form_data aid=1   return 8 items data
	aidStr := c.DefaultPostForm("aid", "0")
	//trans str to int
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid)
	if err != nil{
		ReturnError(c, 4004, "No related info")
	}
	ReturnSuccess(c, 0, "success", rs, 1)

}