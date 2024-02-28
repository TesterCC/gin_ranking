package router

import (
	"gin_ranking/config"
	"gin_ranking/controllers"
	logger "gin_ranking/pkg"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
)

func Router() *gin.Engine{
	r := gin.Default()

	// 以中间件的形式在路由中调用logger
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	// 配置 sessions_redis的信息
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "PenTest123", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// define router directly example
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	user := r.Group("/user")
	// router use controllers   // 大项目推荐用下面这种方式,以防方法名冲突
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)

		// 路径传参 name
		user.GET("/info/:name", controllers.UserController{}.GetUserInfo)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}

	return r
}