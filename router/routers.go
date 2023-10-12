package router

import (
	"gin_ranking/controllers"
	logger "gin_ranking/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine{
	r := gin.Default()

	// 以中间件的形式在路由中调用logger
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	// define router directly example
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	user := r.Group("/user")

	{
		// router use controllers   // 大项目推荐用下面这种方式,以防方法名冲突
		// 路径传参 name
		user.GET("/info/:name", controllers.UserController{}.GetUserInfo)

	}

	return r
}