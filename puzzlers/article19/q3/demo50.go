package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	// right way to use recover
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// wrong way to use recover()
	fmt.Println("no panic: %v\n", recover())

	// raise a panic
	panic(errors.New("something wrong"))

	// wrong way to use recover
	p := recover()
	fmt.Println("panic: %s\n", p)
	fmt.Println("Exit function main.")
}
