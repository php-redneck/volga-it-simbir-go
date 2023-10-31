package httpserver

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Func func(ctx context.Context) error

type Closer struct {
	mutex    sync.Mutex
	funcList []Func
}

func (c *Closer) Add(f Func) {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	c.funcList = append(c.funcList, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	messages := make([]string, 0, len(c.funcList))

	complete := make(chan struct{}, 1)

	go func() {
		for _, f := range c.funcList {
			if err := f(ctx); err != nil {
				messages = append(messages, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}

	if len(messages) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(messages, "\n"),
		)
	}

	return nil
}
