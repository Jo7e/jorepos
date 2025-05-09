// Package chromedp provides a simplified interface to the chromedp browser automation library
package chromedp

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/log"
)

// Browser represents a browser instance with simplified chromedp functionality
type Browser struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger *log.Logger
}

// Options contains configuration options for the browser
type Options struct {
	// Headless determines if the browser runs in headless mode
	Headless bool
	// Timeout is the default timeout for browser operations
	Timeout time.Duration
	// UserAgent is the custom user agent string
	UserAgent string
	// Logger is an optional custom logger
	Logger *log.Logger
	// ExtraOptions contains additional chromedp options
	ExtraOptions []chromedp.ExecAllocatorOption
}

// DefaultOptions returns the default browser options
func DefaultOptions() Options {
	return Options{
		Headless: true,
		Timeout:  30 * time.Second,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}
}

// New creates a new browser instance with the provided options
func New(opts Options) *Browser {
	logger := opts.Logger
	if logger == nil {
		logger = log.New()
	}

	logger.Debug("Initializing browser")

	// Set up the options
	options := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent(opts.UserAgent),
	}

	// Add headless mode if requested
	if opts.Headless {
		options = append(options, chromedp.Headless)
	}

	// Add any extra options
	options = append(options, opts.ExtraOptions...)

	// Create the browser context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, _ := chromedp.NewContext(allocCtx)

	// Create a timeout if specified
	if opts.Timeout > 0 {
		ctx, _ = context.WithTimeout(ctx, opts.Timeout)
	}

	return &Browser{
		ctx:    ctx,
		cancel: cancel,
		logger: logger,
	}
}

// Navigate navigates to the specified URL
func (b *Browser) Navigate(url string) error {
	b.logger.Info("Navigating to URL", "url", url)
	return chromedp.Run(b.ctx, chromedp.Navigate(url))
}

// Screenshot takes a screenshot of the current page
func (b *Browser) Screenshot(selector string) ([]byte, error) {
	var buf []byte
	var err error

	b.logger.Debug("Taking screenshot", "selector", selector)

	if selector == "" {
		// Full page screenshot
		err = chromedp.Run(b.ctx, chromedp.CaptureScreenshot(&buf))
	} else {
		// Element screenshot
		err = chromedp.Run(b.ctx, chromedp.Screenshot(selector, &buf))
	}

	if err != nil {
		b.logger.Error("Failed to take screenshot", "error", err)
		return nil, err
	}

	return buf, nil
}

// Text extracts the text content from an element
func (b *Browser) Text(selector string) (string, error) {
	var text string

	b.logger.Debug("Extracting text", "selector", selector)
	err := chromedp.Run(b.ctx, chromedp.Text(selector, &text))
	if err != nil {
		b.logger.Error("Failed to extract text", "selector", selector, "error", err)
		return "", err
	}

	return text, nil
}

// Click performs a click action on an element
func (b *Browser) Click(selector string) error {
	b.logger.Debug("Clicking element", "selector", selector)
	return chromedp.Run(b.ctx, chromedp.Click(selector))
}

// Type types text into an input field
func (b *Browser) Type(selector, text string) error {
	b.logger.Debug("Typing text", "selector", selector)
	return chromedp.Run(b.ctx, 
		chromedp.Click(selector),
		chromedp.SendKeys(selector, text),
	)
}

// WaitVisible waits for an element to be visible
func (b *Browser) WaitVisible(selector string) error {
	b.logger.Debug("Waiting for element", "selector", selector)
	return chromedp.Run(b.ctx, chromedp.WaitVisible(selector))
}

// Execute runs a custom set of actions
func (b *Browser) Execute(actions ...chromedp.Action) error {
	b.logger.Debug("Executing custom actions")
	return chromedp.Run(b.ctx, actions...)
}

// Close closes the browser
func (b *Browser) Close() {
	b.logger.Debug("Closing browser")
	b.cancel()
}
