package logging

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type LogLevels struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
}

var Logger LogLevels

func init() {
	Logger = initializeLogger()
}

func initializeLogger() LogLevels {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv("LOG_PATH") == "" {
		log.Fatal("LOG_PATH environment variable is not set")
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

	logLevels := LogLevels{
		Info:    log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Warning: log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error:   log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		Debug:   log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return logLevels
}
