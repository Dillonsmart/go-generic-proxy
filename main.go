package main

import (
	"github.com/dillonsmart/go-generic-proxy/logging"
	"github.com/gin-gonic/gin"
)

func main() {
	logging.Logger.Info.Println("Starting the server...")

	r := gin.Default()
	RegisterRoutes(r)
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
