package cdp

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/log"
)

// PasteInForm pastes a value into a form element
func (b *Browser) PasteInForm(inputSelector, value string) error {
	if err := chromedp.Run(b.ctx,
		chromedp.WaitVisible(inputSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Debug("Paste-ing", "element", inputSelector)
			return b.setValue(inputSelector, value)
		}),
		chromedp.Sleep(time.Duration(rand.Intn(400))),
	); err != nil {
		return err
	}

	return nil
}

// setValue sets the value of an input element
func (b *Browser) setValue(inputSelector, value string) error {
	if err := chromedp.Run(b.ctx,
		chromedp.SetValue(inputSelector, value, chromedp.ByQuery),
	); err != nil {
		return err
	}

	return nil
}
