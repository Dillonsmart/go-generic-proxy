package main

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	// do stuff
	r := gin.Default()
	RegisterRoutes(r)
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
