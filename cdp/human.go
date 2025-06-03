package cdp

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jo7e/jorepos/log"
)

func RandomDelay(min, max int) time.Duration {
	delay := min + rand.Intn(max-min)
	log.Debug("", "delay", delay)
	return time.Duration(delay) * time.Millisecond
}

var ErrIncorrectParametersNumber = errors.New("pass 0 or 2 ints for this function")

func (b *Browser) HumanTyping(selector, text string, delays ...int) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var typingMinDelay, typingMaxDelay int
		switch len(delays) {
		case 0:
			typingMinDelay = b.TypingWaitMinTime
			typingMaxDelay = b.TypingWaitMaxTime
		case 2:
			typingMinDelay = delays[0]
			typingMaxDelay = delays[1]
		default:
			return fmt.Errorf("invalid delay parameters: %w", ErrIncorrectParametersNumber)
		}

		if err := chromedp.Run(b.ctx,
			chromedp.ActionFunc(func(ctx context.Context) error {
				for _, char := range text {
					if err := chromedp.Run(b.ctx,
						chromedp.SendKeys(selector, string(char), chromedp.ByQuery),
						chromedp.Sleep(RandomDelay(typingMinDelay, typingMaxDelay)),
					); err != nil {
						return err
					}
				}
				return nil
			}),
		); err != nil {
			return err
		}
		return nil
	})
}

func (b *Browser) HumanAfterAction(delays ...int) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var minDelay, maxDelay int

		switch len(delays) {
		case 0:
			minDelay = b.AfterActionMinWaitTime
			maxDelay = b.AfterActionMaxWaitTime
		case 2:
			minDelay = delays[0]
			maxDelay = delays[1]
		default:
			return fmt.Errorf("invalid delay parameters: %w", ErrIncorrectParametersNumber)
		}

		if err := chromedp.Run(b.ctx,
			chromedp.Sleep(RandomDelay(minDelay, maxDelay)),
		); err != nil {
			return err
		}
		return nil
	})
}
