package main

import (
	"example/web-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
}
