package main

import (
	"fmt"
	"github.com/dillonsmart/go-generic-proxy/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

func RegisterRoutes(router *gin.Engine) {
	initLogger(router)
	router.Use(gin.Recovery())

	router.GET("/ping", controllers.Ping)
	router.Any("/proxy/*path", controllers.HandleAny)
}

func initLogger(router *gin.Engine) {
	envError := godotenv.Load()

	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	date := time.Now().Format("2006-01-02")
	workingDir, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal("Error getting working directory: ", wdErr)
	}
	logDir := workingDir + "/" + os.Getenv("LOG_PATH")
	logFileName := logDir + "/log-" + date + ".log"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatal("Error creating log directory: ", err)
		}
	}

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
}
