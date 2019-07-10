package main

import (
	"log"
	"sync"
	"time"
)

// 计数器
type counter struct {
	num uint         // 计数
	mu  sync.RWMutex // 读写锁
}

// 返回当前计数，加锁读。
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

// 增加计数
func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func count(c *counter) {
	sign := make(chan struct{}, 3)
	// 用于累加计数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 1, j)
		}
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for k := 1; k <= 20; k++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number in counter: %d [%d-%d]",
				c.number(), 2, k)

		}
	}()
	<-sign
	<-sign
	<-sign
}

func main() {
	c := counter{}
	count(&c)
}
