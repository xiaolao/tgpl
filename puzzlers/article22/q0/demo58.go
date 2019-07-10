package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

// protecting 用于指示是否使用互斥锁来保护数据写入。
// 若值等于0则表示不使用，若值大于0则表示使用。
// 改变该变量的值，然后多运行几次程序，并观察程序打印的内容。
var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1,
		"It indicates whether to use a mutex to protect writing")
}

func main() {
	// 解析命令行参数
	flag.Parse()
	// buffer代表缓冲区
	var buffer bytes.Buffer

	const (
		max1 = 5  // 代表启动的goroutine的数量
		max2 = 10 // 代表每个goroutine需要写入的数据块的数量
		max3 = 10 // 代表每个数据块中需要有多少个重复的数字
	)

	// mu 代表流程要使用的互斥锁
	var mu sync.Mutex
	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= max2; j++ {
				// 准备数据
				header := fmt.Sprintf("\n[id: %d, iteration: %d]", id, j)
				data := fmt.Sprintf(" %d", id*j)
				// 写入数据
				if protecting > 0 {
					mu.Lock()
				}
				// 写入头信息
				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error: %s [%d]", err, id)
				}
				// 写入10遍data
				for k := 0; k < max3; k++ {
					_, err := writer.Write([]byte(data))
					if err != nil {
						log.Printf("error: %s [%d]", err, id)
					}
				}
				if protecting > 0 {
					mu.Unlock()
				}
			}
		}(i, &buffer)
	}
	// 阻塞等待前面的goroutine结束
	for i := 0; i < max1; i++ {
		<-sign
	}
	// 读取出所有byte字节
	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	log.Printf("The contents:\n%s", data)
}
