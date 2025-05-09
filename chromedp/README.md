# ChromeDP Package

A simplified wrapper around the [chromedp](https://github.com/chromedp/chromedp) browser automation library for Go.

## Features

- Simple, fluent interface for common browser automation tasks
- Integrated logging with the jorepos/log package
- Sensible defaults with customizable options
- Helper methods for common tasks: navigation, screenshots, text extraction, clicking, typing

## Usage

```go
package main

import (
	"context"
	"os"
	"time"

	"github.com/jo7e/jorepos/chromedp"
	"github.com/jo7e/jorepos/log"
)

func main() {
	// Create a new browser with default options (headless mode)
	browser := chromedp.New(chromedp.DefaultOptions())
	defer browser.Close()

	// Navigate to a website
	err := browser.Navigate("https://example.com")
	if err != nil {
		log.Fatal("Failed to navigate", "error", err)
	}

	// Get text from a selector
	text, err := browser.Text("h1")
	if err != nil {
		log.Fatal("Failed to get text", "error", err)
	}
	log.Info("Found text", "text", text)

	// Take a screenshot
	screenshot, err := browser.Screenshot("")
	if err != nil {
		log.Fatal("Failed to take screenshot", "error", err)
	}

	// Save the screenshot
	err = os.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		log.Fatal("Failed to save screenshot", "error", err)
	}
	log.Info("Screenshot saved", "path", "screenshot.png")
}
```

### Custom Options

```go
package main

import (
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/chromedp"
	"github.com/jo7e/jorepos/log"
)

func main() {
	// Create a custom logger
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	// Create a browser with custom options
	browser := chromedp.New(chromedp.Options{
		Headless:  false,                // Show the browser
		Timeout:   60 * time.Second,     // Longer timeout
		UserAgent: "Custom User Agent",  // Custom user agent
		Logger:    logger,               // Custom logger
		ExtraOptions: []chromedp.ExecAllocatorOption{
			chromedp.WindowSize(1920, 1080),
			chromedp.NoSandbox, 
		},
	})
	defer browser.Close()

	// Fill a form
	_ = browser.Navigate("https://example.com/login")
	_ = browser.Type("#username", "user@example.com")
	_ = browser.Type("#password", "password")
	_ = browser.Click("#login-button")

	// Wait for dashboard to load
	_ = browser.WaitVisible("#dashboard")
	log.Info("Successfully logged in")
}
```
