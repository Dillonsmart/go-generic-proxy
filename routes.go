package main

import (
	"github.com/dillonsmart/go-generic-proxy/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", controllers.Ping)
	router.Any("/proxy/*path", controllers.HandleAny)
}
