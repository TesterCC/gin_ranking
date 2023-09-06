package router

import (
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
	return r
}