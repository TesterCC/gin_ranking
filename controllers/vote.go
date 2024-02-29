package controllers

import (
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// call all kind of models method to get data

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	// get user id (voter id)  candidate id
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "Please input correct info")
		return
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "voter isn't exist")
		return
	}

	play, _ := models.GetPlayerIno(playerId)
	if play.Id == 0 {
		ReturnError(c, 4002, "player isn't exist")
		return
	}

	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4003, "already voted")
		return
	}

	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		// update vote score info, score++
		models.UpdatePlayerScore(playerId)
		ReturnSuccess(c, 0, "vote success", rs, 1)
		return
	}
    ReturnError(c, 4004, "unknown error, please contract admin")
	return
}
