# Log Package

A customized wrapper around [charmbracelet/log](https://github.com/charmbracelet/log) with sensible defaults, environment variable configuration, and custom styling.

## Features

- Pre-configured with reasonable defaults
- Environment variable configuration
- Custom styling with [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)
- Built-in support for structured logging

## Configuration

The logger can be configured using the following environment variables:

```
LOG_LEVEL=debug|info|warn|error  # Sets the log level (defaults to info)
LOG_TIMESTAMP=true               # Enables timestamp reporting
LOG_CALLER=true                  # Enables caller reporting
```

## Usage

```go
package main

import (
	"os"

	"github.com/jo7e/jorepos/log"
)

func main() {
	// Initialize the default logger with environment variable configuration
	log.DefaultLogger()
	
	// Use the configured default logger
	log.Info("Hello, world!")
	
	// With fields
	log.With("user", "jo7e").With("action", "login").Info("User logged in")
	
	// Debug logs (only visible when level is Debug or lower)
	log.Debug("Debug message")
	
	// Error logs
	log.Error("Something went wrong")
}
```

## Custom Styling

The package provides custom styling for logs using lipgloss:

```go
package main

import (
	"os"
	"github.com/jo7e/jorepos/log"
)

func main() {
	// Create a custom logger
	logger := log.New(os.Stdout)
	
	// Apply custom styles
	logger = log.ApplyStyles(logger)
	
	logger.Info("This message will use custom styling")
}
```

The styling includes custom colors for prefixes and various log elements to improve readability.
