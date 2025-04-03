package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	RegisterRoutes(r)
	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
