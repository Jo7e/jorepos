package log

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

type Logger log.Logger

// DefaultLogger configures the default logger with custom styles and options
func DefaultLogger() {
	logLevel, err := log.ParseLevel(strings.ToLower(os.Getenv("LOG_LEVEL")))
	if err != nil {
		logLevel = log.InfoLevel
	}
	logger := log.New(os.Stdout)

	logger.SetLevel(logLevel)

	if os.Getenv("LOG_TIMESTAMP") == "true" {
		logger.SetReportTimestamp(true)
	}

	if os.Getenv("LOG_CALLER") == "true" {
		logger.SetReportCaller(true)
	}

	logger = ApplyStyles(logger)
	log.SetDefault(logger)
}


func Info(msg string, keysAndValues ...interface{}) {
	log.Info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	log.Warn(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	log.Debug(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	log.Error(msg, keysAndValues...)
}
	