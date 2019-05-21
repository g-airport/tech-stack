package currency_pattern

import (
	"context"
	"testing"
	"time"
)

func BenchmarkUseTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := make(chan Msg, 1)
		msg := Msg{data: 1}
		UseTimer(time.Millisecond, msg, c)
	}
}

func BenchmarkUseContext(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := make(chan Msg, 1)
		msg := Msg{data: 1}
		ctx := context.Background()
		UseContext(ctx, time.Millisecond, msg, c)
	}

}

func BenchmarkUseFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 := make(chan Msg, 1)
		r2 := make(chan Msg, 1)
		msg := Msg{data: 1}
		UseFor(msg, r1, r2)
	}
}

func BenchmarkUseWaitGroup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r1 := make(chan Msg, 1)
		r2 := make(chan Msg, 1)
		msg := Msg{data: 1}
		UseWaitGroup(msg, r1, r2)
	}
}

func BenchmarkUseSliceChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chs := make(chan Msg, 2)
		msg := Msg{data: 1}
		UseSliceChan(msg, chs)
	}
}

func BenchmarkUseReflectChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chs := make(chan Msg, 2)
		msg := Msg{data: 1}
		UseReflectChan(msg, chs)
	}
}

func BenchmarkUseTimerSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := Msg{data: 1}
		r1 := make(chan Msg, 1)
		r2 := make(chan Msg, 1)
		UseTimerSelect(time.Millisecond, msg, r1, r2)
	}
}

func BenchmarkUseContextGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msg := Msg{data: 1}
		r1 := make(chan Msg, 1)
		r2 := make(chan Msg, 1)
		ctx := context.Background()
		UseContextGoroutine(ctx, time.Millisecond, msg, r1, r2)
	}
}
