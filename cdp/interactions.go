package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/log"
)

// Click performs a click action on an element
func (b *Browser) Click(buttonSelector string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		log.Debug("Clicking", "element", buttonSelector)
		if err := chromedp.Run(b.ctx,
			chromedp.WaitVisible(buttonSelector, chromedp.ByQuery),
			chromedp.Click(buttonSelector, chromedp.ByQuery),
			b.HumanAfterAction(),
		); err != nil {
			return err
		}

		return nil
	})
}

// Type types text into an input field
func (b *Browser) Type(selector, text string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		log.Debug("Typing text", "selector", selector)
		if err := chromedp.Run(b.ctx,
			b.HumanTyping(selector, text),
			b.HumanAfterAction(),
		); err != nil {
			return err
		}
		return nil
	})
}

// ClickAndType clicks on an element and types text into it
func (b *Browser) ClickAndType(buttonSelector, text string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		if err := chromedp.Run(b.ctx,
			b.Click(buttonSelector),
			b.Type(buttonSelector, text),
		); err != nil {
			return err
		}
		return nil
	})
}
