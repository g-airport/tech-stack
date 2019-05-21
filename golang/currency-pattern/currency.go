package currency_pattern

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// message 数据
type Msg struct {
	data int
}

// is done 是否处理完成
type Done bool

func UseTimer(d time.Duration, msg Msg, c chan Msg) Done {
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case c <- msg:
		return true
	case <-t.C:
		return false
	}
}

func UseContext(ctx context.Context, d time.Duration, msg Msg, c chan Msg) Done {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	select {
	case c <- msg:
		return true
	case <-ctx.Done():
		return false
	}
}

// First come first serve
func UseFor(msg Msg, r1, r2 chan<- Msg) {
	for i := 0; i < 2; i++ {
		select {
		case r1 <- msg:
			r1 = nil
		case r2 <- msg:
			r2 = nil

		}
	}
}

func UseWaitGroup(msg Msg, r1, r2 chan<- Msg) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { r1 <- msg; wg.Done() }()
	go func() { r2 <- msg; wg.Done() }()
}

//
func UseSliceChan(msg Msg, chs ...chan<- Msg) {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, c := range chs {
		// copy
		c := c
		go func() { c <- msg; wg.Done() }()
	}
}

func UseReflectChan(msg Msg, chs ...chan<- Msg) {
	cases := make([]reflect.SelectCase, len(chs))
	for i, ch := range chs {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(msg),
		}
	}
	for i := 0; i < len(chs); i++ {
		chosen, _, _ := reflect.Select(cases)
		cases[chosen].Chan = reflect.ValueOf(nil)

	}
}

// timer && select
func UseTimerSelect(d time.Duration, msg Msg, r1, r2 chan Msg) int {
	t := time.NewTimer(d)
	for i := 0; i < 2; i++ {
		select {
		case r1 <- msg:
			r1 = nil
		case r2 <- msg:
			r2 = nil

		case <-t.C:
			return i

		}
	}
	t.Stop()
	return 2
}

// context && goroutine
func UseContextGoroutine(ctx context.Context, d time.Duration, msg Msg, ch ...chan Msg) int {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	var (
		wr int32
		wg sync.WaitGroup
	)
	wg.Add(len(ch))
	for _, c := range ch {
		c := c
		go func() {
			defer wg.Done()
			select {
			case c <- msg:
				atomic.AddInt32(&wr, 1)
			case <-ctx.Done():

			}
		}()
	}
	wg.Wait()
	return int(wr)
}
