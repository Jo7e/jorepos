// Package log provides a customized logging implementation based on charmbracelet/log
package log

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

// Logger is a wrapper around charmbracelet/log Logger
type Logger struct {
	*log.Logger
}

// Default logger instance
var defaultLogger = New()

// New creates a new Logger instance with default configuration
func New() *Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Prefix:          "jorepos",
	})
	return &Logger{
		Logger: logger,
	}
}

// WithOutput creates a new logger with a custom output writer
func WithOutput(w io.Writer) *Logger {
	l := New()
	l.SetOutput(w)
	return l
}

// SetLevel sets the logging level of the default logger
func SetLevel(level log.Level) {
	defaultLogger.SetLevel(level)
}

// Debug logs a debug message using the default logger
func Debug(msg interface{}, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}

// Info logs an info message using the default logger
func Info(msg interface{}, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}

// Warn logs a warning message using the default logger
func Warn(msg interface{}, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}

// Error logs an error message using the default logger
func Error(msg interface{}, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}

// Fatal logs a fatal message using the default logger and exits
func Fatal(msg interface{}, args ...interface{}) {
	defaultLogger.Fatal(msg, args...)
}

// WithFields creates a new entry with the given fields
func WithFields(fields log.Fields) *log.Logger {
	return defaultLogger.With(fields)
}
