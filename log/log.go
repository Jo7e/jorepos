package log

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

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
