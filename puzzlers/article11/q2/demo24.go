package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

func example1() {
	// int a array which contain 3 channels
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// choose one channel randomly and send a element to it.
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	//
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Println("No candidate case is selected!")
	}
}

func example2() {
	intChan := make(chan int, 1)
	// close the channel after one second.
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
