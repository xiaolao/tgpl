package main

import "fmt"

func main() {

	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
	// 1,2,3,4,5,6,7   7, 7

	s9 := a1[1:4]
	fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	// 2,3,4    3, 6

	for i := 1; i <= 5; i++ {
		s9 := append(s9, i)
		fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))
	}
	// 2,3,4,1          4, 6
	// 2,3,4,1,2        5, 6
	// 2,3,4,1,2,3      6, 6
	// 2,3,4,1,2,3,4    7, 12
	// 2,3,4,1,2,3,4,5  8, 12

	fmt.Printf("a1: %v (len: %d, cap: %d)\n", a1, len(a1), cap(a1))
	// 1,2,3,4,1,2,3   7, 7
	fmt.Println()
}
