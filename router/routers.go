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
		// router use controllers   // 小项目接口少可以不用结构体方法调用，大项目推荐用下面这种方式,以防方法名冲突
		// 参数传递1 直接拼接
		user.GET("/info/:name", controllers.UserController{}.GetUserInfo)
		user.GET("/information", controllers.UserController{}.GetUserInformation)  // custom dev
		user.POST("/list", controllers.UserController{}.GetList)   // test exception catch
		user.POST("/list/test", controllers.UserController{}.GetUserListTest)

		user.POST("/add", controllers.UserController{}.AddUser)
		//user.POST("/add", func(ctx *gin.Context) {
		//	//ctx.String(http.StatusOK, "user add")
		//	ctx.JSON(http.StatusOK, "user add")
		//})
		user.PUT("/edit", controllers.UserController{}.UpdateUser)

		//user.PUT("/edit", func(ctx *gin.Context) {
		//	ctx.String(http.StatusOK, "user edit")
		//})

		user.DELETE("/delete", controllers.UserController{}.DeleteUser)
	}

	order := r.Group("/order")

	{
		order.GET("/info", controllers.OrderController{}.GetUserInfo)
		order.POST("/list", controllers.OrderController{}.GetList)
		order.POST("/listAll", controllers.OrderController{}.GetListAll)
		order.POST("/listAllSearch", controllers.OrderController{}.GetListAllSearch)
	}

	return r
}