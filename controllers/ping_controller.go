package controllers

import (
	"github.com/dillonsmart/go-generic-proxy/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	logging.Logger.Info.Println("Request to /ping received")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
