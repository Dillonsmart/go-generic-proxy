package controllers

import (
	"github.com/dillonsmart/go-generic-proxy/logging"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func HandleAny(c *gin.Context) {
	logging.Logger.Info.Println("Request to /proxy/" + c.Param("path") + " received")
	c.JSON(200, gin.H{
		"message": "Forwarding request received",
		"path":    c.Param("path"),
		"method":  c.Request.Method,
		"headers": c.Request.Header,
	})

	forwardRequest(c)
}

func forwardRequest(c *gin.Context) {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	forwardTo := os.Getenv("FORWARD_TO")
	method := c.Request.Method

	if forwardTo == "" {
		log.Println("FORWARD_TO environment variable is not set")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	path := strings.Replace(c.Param("path"), "/proxy", "", -1)
	url := forwardTo + path
	println("forwarding request to: " + url)

	makeRequest(c, url, method)
}

func makeRequest(c *gin.Context, url string, method string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating GET request: %v", err)
	}

	// copies the headers from the original request to the new request
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error forwarding GET request: %v", err)
	}
	defer resp.Body.Close()

	// copies response headers from the forwarded request to the original response
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(resp.StatusCode)
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		log.Fatalf("Error writing response body: %v", err)
	}
}
