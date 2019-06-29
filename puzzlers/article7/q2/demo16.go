package main

import "fmt"

func main() {

	// 当原slice的capaticy小于1024时,扩容后的cap为2*capaticy
	s6 := make([]int, 0)
	fmt.Printf("The capacity of s6: %d\n", cap(s6)) // len(s6) == 0, cap(s6) == 0,
	for i := 1; i <= 5; i++ {
		s6 = append(s6, i)
		fmt.Printf("s6(%d): %v, len: %d, cap: %d\n", i, s6, len(s6), cap(s6))
	}
	// s6(1): len: 1, cap: 1
	// s6(2): len: 2, cap: 2
	// s6(3): len: 3, cap: 4
	// s6(4): len: 4, cap: 4
	// s6(5): len: 5, cap: 8
	fmt.Println()

	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7)) // 1024
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1)) // 1224, 1024*1.25
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2)) // 1224, 1024*1.25
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3)) // 1224, 1024*1.25

	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8)) // 10
	s8a := append(s8, make([]int, 11)...)
	fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a)) // 21, 22
	s8b := append(s8a, make([]int, 23)...)
	fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b)) // 44, 44
	s8c := append(s8b, make([]int, 45)...)
	fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c)) // 89, 96

}
