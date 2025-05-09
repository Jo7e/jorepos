# Log Package

A customized wrapper around [charmbracelet/log](https://github.com/charmbracelet/log) with sensible defaults and additional convenience methods.

## Features

- Pre-configured with reasonable defaults
- Simple interface that matches the standard library's log package
- Built-in support for structured logging
- Pretty output formatting using charmbracelet's styling

## Usage

```go
package main

import (
	"github.com/jo7e/jorepos/log"
)

func main() {
	// Use default logger
	log.Info("Hello, world!")
	
	// With fields
	log.WithFields(log.Fields{
		"user": "jo7e",
		"action": "login",
	}).Info("User logged in")
	
	// Change log level
	log.SetLevel(log.DebugLevel)
	
	// Debug logs (only visible when level is Debug or lower)
	log.Debug("Debug message")
	
	// Error logs
	log.Error("Something went wrong")
}
```

## Custom Logger

```go
package main

import (
	"os"
	
	"github.com/jo7e/jorepos/log"
)

func main() {
	// Create a custom logger with file output
	f, _ := os.Create("app.log")
	logger := log.WithOutput(f)
	
	logger.Info("This will be written to app.log")
}
```
