package cdp

import (
	"context"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/log"
)

// Click performs a click action on an element
func (b *Browser) Click(buttonSelector string) error {
	if err := chromedp.Run(b.ctx,
		chromedp.WaitVisible(buttonSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Debug("Clicking", "element", buttonSelector)
			return nil
		}),
		chromedp.Click(buttonSelector, chromedp.ByQuery),
		chromedp.Sleep(time.Duration(rand.Intn(400))),
	); err != nil {
		return err
	}

	return nil
}
