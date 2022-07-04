package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()

	v1API := router.Group("/v1/api")
	{
		v1API.GET("/todo", GetTodo)
		v1API.GET("/ping", GetPing)

		v1API.POST("/todo", PostTodo)
	}

	return router
}
