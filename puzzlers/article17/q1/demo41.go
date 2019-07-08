package main

import "fmt"

func main() {
	// example 1: 当for语句执行的时候,在range关键字右边的
	// numbers1会先求值，由于slice是引用类型，故range语句不
	// 拷贝副本,slice的迭代变量有两个，这里只有一个。
	// 当只有一个迭代变量的时候，数组、数组的指针、切片
	// 和字符串的元素值都是无处安放的，我们只能拿到按照从小到大下标值。
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			// |= 按位或 1001 |= 1011 -> 1011
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1) // [1, 2, 3, 7, 5, 6]
	fmt.Println()

	// example 2: 数组为值类型，跟在range后面会有拷贝
	// for迭代的其实是数组的副本
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2) // [7, 3, 5, 7, 9, 11]
	fmt.Println()

	// example3: slice为引用类型不会被拷贝
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3) // [22, 3, 6, 10, 15, 21]
}
