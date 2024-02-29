package controllers

import (
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	// http://localhost:9999/player/list post form_data aid=1   return 8 items data
	aidStr := c.DefaultPostForm("aid", "0")
	//trans str to int
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc") // 根据id升序排序 id asc ，按id降序 id desc
	if err != nil {
		ReturnError(c, 4004, "No related info")
	}
	ReturnSuccess(c, 0, "success", rs, 1)

}

// GetRanking mysql orderby 排序实现排名展示
func (p PlayerController) GetRanking(c *gin.Context) {

	//err := cache.Rdb.Set(cache.Rctx, "name", "zstest", 0).Err()
	//if err != nil {
	//	panic(err)
	//}

	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)   // activity id

	// https://www.imooc.com/video/24568  5-2 05:30
	// todo 新增逻辑：先从redis中获取数据，当redis中数据不存在时，再从mysql数据库中获取数据，且从mysql中获取数据后要存入redis中
	// 使用redis时要注意设置过期时间

	// method 1
	//var redisKey string
	//redisKey = "ranking:" + aidStr

	// method 2 simplify
	redisKey := "ranking:" + aidStr

	rs, err := models.GetPlayers(aid, "score desc") // 按score字段降序排序
	if err != nil {
		ReturnError(c, 4004, "No data")
		return
	}

	ReturnSuccess(c, 0, "request success", rs, 1)
	return
}
