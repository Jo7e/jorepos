// Package cdp provides a simplified interface to the chromedp browser automation library
package cdp

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
}

type ExtraOptions []chromedp.ExecAllocatorOption

var DefaultExecAllocatorOptions = chromedp.DefaultExecAllocatorOptions

// Options contains configuration options for the browser
type Options struct {
	// Headless determines if the browser runs in headless mode
	Headless bool
	// Timeout is the default timeout for browser operations
	Timeout time.Duration
	// UserAgent is the custom user agent string
	UserAgent string
	// ExtraOptions contains additional chromedp options
	ExtraOptions ExtraOptions
}

// DefaultOptions returns the default browser options
func DefaultOptions() Options {
	return Options{
		Headless:  true,
		Timeout:   30 * time.Second,
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}
}

// New creates a new browser instance with the provided options
func New(opts Options) *Browser {
	log.Debug("Initializing browser")

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
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), options...)
	ctx, cancel := chromedp.NewContext(allocCtx)

	// Create a timeout if specified
	if opts.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, opts.Timeout)
	}

	return &Browser{
		ctx: ctx,
		cancel: func() {
			cancel()
			allocCancel()
		},
	}
}

// Flag creates an allocator option for a command line flag
func Flag(name string, value any) chromedp.ExecAllocatorOption {
	return chromedp.Flag(name, value)
}

// Sleep creates an action to pause for a specified duration
func (b *Browser) Sleep(duration time.Duration) error {
	log.Debug("Sleeping for duration", "duration", duration)
	return chromedp.Run(b.ctx, chromedp.Sleep(duration))
}

// Run executes a set of chromedp actions
func (b *Browser) Run(actions ...chromedp.Action) error {
	return chromedp.Run(b.ctx, actions...)
}

// Navigate navigates to the specified URL
func (b *Browser) Navigate(url string) error {
	log.Info("Navigating to URL", "url", url)
	return chromedp.Run(b.ctx, chromedp.Navigate(url))
}

// Screenshot takes a screenshot of the current page
func (b *Browser) Screenshot(selector string) ([]byte, error) {
	var buf []byte
	var err error

	log.Debug("Taking screenshot", "selector", selector)

	if selector == "" {
		// Full page screenshot
		err = chromedp.Run(b.ctx, chromedp.CaptureScreenshot(&buf))
	} else {
		// Element screenshot
		err = chromedp.Run(b.ctx, chromedp.Screenshot(selector, &buf))
	}

	if err != nil {
		log.Error("Failed to take screenshot", "error", err)
		return nil, err
	}

	return buf, nil
}

// Text extracts the text content from an element
func (b *Browser) Text(selector string) (string, error) {
	var text string

	log.Debug("Extracting text", "selector", selector)
	err := chromedp.Run(b.ctx, chromedp.Text(selector, &text))
	if err != nil {
		log.Error("Failed to extract text", "selector", selector, "error", err)
		return "", err
	}

	return text, nil
}

// Click performs a click action on an element
func (b *Browser) Click(selector string) error {
	log.Debug("Clicking element", "selector", selector)
	return chromedp.Run(b.ctx, chromedp.Click(selector))
}

// Type types text into an input field
func (b *Browser) Type(selector, text string) error {
	log.Debug("Typing text", "selector", selector)
	return chromedp.Run(b.ctx,
		chromedp.Click(selector),
		chromedp.SendKeys(selector, text),
	)
}

// WaitVisible waits for an element to be visible
func (b *Browser) WaitVisible(selector string) error {
	log.Debug("Waiting for element", "selector", selector)
	return chromedp.Run(b.ctx, chromedp.WaitVisible(selector))
}

// Execute runs a custom set of actions
func (b *Browser) Execute(actions ...chromedp.Action) error {
	log.Debug("Executing custom actions")
	return chromedp.Run(b.ctx, actions...)
}

// Close closes the browser
func (b *Browser) Close() {
	log.Debug("Closing browser")
	b.cancel()
}
