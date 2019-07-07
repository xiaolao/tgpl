package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// 通过使用全局变量count来判断是否执行匿名函数fn
	// 从而达到控制打印输出顺序的结果。
	var count uint32
	trigger := func(i uint32, fn func()) {
		for {
			// LoadUint32 atomically loads *addr.
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				// AddUint32 atomically adds delta to *addr
				// and returns the new value.
				// To subtract a signed positive constant value c from x,
				// do AddUint32(&x, ^uint32(c-1)).
				// In particular, to decrement x, do AddUint32(&x, ^uint32(0)).
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}

	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	// 阻塞主goroutine
	trigger(10, func() {})
}
