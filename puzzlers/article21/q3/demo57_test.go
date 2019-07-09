package q3

import (
	"testing"
	"time"
)

func BenchmarkGetPrimes(b *testing.B) {
	b.StopTimer()
	time.Sleep(time.Millisecond * 500)
	max := 10000
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GetPrimes(max)
	}
}
