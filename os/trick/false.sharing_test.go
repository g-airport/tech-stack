package trick

import (
	"sync"
	"sync/atomic"
	"testing"

)

var num = 100000

type PaddedUint64 struct {
	_p1,_p2,_p3,_p4,_p5,_p6,_p7 uint64
	value uint64
}

func NewPaddedUint64(v uint64) PaddedUint64 {
	return PaddedUint64{value:v}
}

func BenchmarkType(b *testing.B) {
	values := make([]uint64, 100)
	var start uint64 = 0
	var end uint64 = 0
	var wg sync.WaitGroup

	b.StartTimer()
	wg.Add(num)
	go func() {
		for i := 0; i < num; i++ {
			go func() {
				t := atomic.AddUint64(&end, 1)
				values[t%100] = 1
			}()
		}
	}()
	go func() {
		for i := 0; i < num; i++ {
			go func() {
				t := atomic.AddUint64(&start, 1)
				values[t%100] = 1
				wg.Done()

			}()
		}
	}()
	wg.Wait()
	b.StopTimer()
}

func BenchmarkPaddedType(b *testing.B) {
	values := make([]uint64, 100)
	var start = NewPaddedUint64(0)
	var end = NewPaddedUint64(0)
	startV := NewPaddedUint64(0).value
	endV := NewPaddedUint64(1).value
	var wg sync.WaitGroup

	b.StartTimer()
	wg.Add(num)
	go func() {
		for i := 0; i < num; i++ {
			go func() {
				t := atomic.AddUint64(&end.value, 1)
				values[t%100] = endV
			}()
		}
	}()
	go func() {
		for i := 0; i < num; i++ {
			go func() {
				t := atomic.AddUint64(&start.value, 1)
				values[t%100] = startV
				wg.Done()

			}()
		}
	}()
	wg.Wait()
	b.StopTimer()
}
