package controllers

import "github.com/gin-gonic/gin"

type VoteController struct {}

func (v VoteController) AddVote(c *gin.Context)  {
     // get user id (voter id)  candidate id
	userIdStr := c.DefaultPostForm("userId", "0")
    // https://www.imooc.com/video/24565  8:00
}
