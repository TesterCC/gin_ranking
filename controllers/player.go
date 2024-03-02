package controllers

import (
	"gin_ranking/cache"
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
	aid, _ := strconv.Atoi(aidStr) // activity id

	// 新增逻辑：先从redis中获取数据，当redis中数据不存在时，再从mysql数据库中获取数据，且从mysql中获取数据后要存入redis中
	// 使用redis时要注意设置过期时间

	// method 1
	//var redisKey string
	//redisKey = "ranking:" + aidStr

	// method 2 simplify
	redisKey := "ranking:" + aidStr

	// 通过有序集合获取 https://pkg.go.dev/github.com/redis/go-redis/v9#Client.ZRevRange
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()

	if err == nil && len(rs) > 0 {
		return
	}

	rsDb, errDb := models.GetPlayers(aid, "score desc") // 按score字段降序排序
	if errDb == nil {
		for _, value := range rsDb {
			// add data in redis, members need define a function in redis.go
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		// set expire time       24 * time.Hour
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Second)  // todo test
		// https://www.imooc.com/video/24568  5-2 10:26
		ReturnSuccess(c, 0, "request success", rs, 1)
		return
	}


	ReturnError(c, 4004, "No data")
	return
}
