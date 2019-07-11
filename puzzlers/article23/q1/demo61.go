package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// mailbox 代表信箱
	// 0表示信箱是空的，1代表信箱是满的。
	var mailbox uint8
	// lock代表信箱上的锁
	var lock sync.RWMutex
	// sendCond 代表专用于发信的条件变量。
	sendCond := sync.NewCond(&lock)
	// recvCond 代表专用于收信的条件变量。
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 3)
	max := 5
	// 发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			// 这边是for语句
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)
	// 用于收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]: the mailbox is full.", j)
			mailbox = 0
			log.Printf("receiver [%d]: the letter has been receiverd.", j)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
