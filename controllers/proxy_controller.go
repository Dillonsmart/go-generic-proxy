package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func HandleAny(c *gin.Context) {
	// Handle any request here
	c.JSON(200, gin.H{
		"message": "Forwarding request received",
		"path":    c.Param("path"),
		"method":  c.Request.Method,
		"headers": c.Request.Header,
	})

	forwardRequest(c)
	// TODO - forward the request to another server
}

func forwardRequest(c *gin.Context) {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	forwardTo := os.Getenv("FORWARD_TO")

	if forwardTo == "" {
		log.Fatalf("FORWARD_TO environment variable is not set")
	}

	// TODO - Add any authorization headers or other headers to the request if needed

	println("forwarding request to: " + forwardTo)

}
