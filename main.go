package main

import (
	"gin_ranking/router"
)

// curl http://127.0.0.1:9999/hello

func main() {

	//// catch exception
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常:", err)
	//	}
	//}()

	//// defer recover panic // 先defer的后执行，后defer的先执行
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	//// cause crash
	//panic("000")   // 3 2 1  捕获异常： 000

	r := router.Router()
	r.Run(":9999")




	//r := gin.Default()
	//
	//// define router directly example
	//r.GET("/hello", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "Hello World!")
	//})

	//r.POST("/user/add", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "user add")
	//})
	//
	//r.PUT("/user/edit", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "user edit")
	//})
	//
	//r.DELETE("/user/delete", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "user delete")
	//
	//})

	//// gin的router支持分组
	//user := r.Group("/user")
	//
	//{
	//	user.POST("/add", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "user add")
	//	})
	//
	//	user.PUT("/edit", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "user edit")
	//	})
	//
	//	user.DELETE("/delete", func(ctx *gin.Context) {
	//		ctx.String(http.StatusOK, "user delete")
	//
	//	})
	//}

}
