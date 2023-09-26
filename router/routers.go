package router

import (
	"gin_ranking/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine{
	r := gin.Default()


	// define router directly example
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	user := r.Group("/user")

	{
		// router use controllers   // 小项目接口少可以不用结构体方法调用，大项目推荐用下面这种方式
		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)

		user.POST("/add", func(ctx *gin.Context) {
			//ctx.String(http.StatusOK, "user add")
			ctx.JSON(http.StatusOK, "user add")
		})

		user.PUT("/edit", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user edit")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")

		})
	}

	order := r.Group("/order")

	{
		order.GET("/info", controllers.OrderController{}.GetUserInfo)
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}