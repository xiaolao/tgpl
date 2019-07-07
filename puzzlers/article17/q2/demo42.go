package main

import "fmt"

func main() {
	// example 1
	// value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// 1 + 3为int类型,与case表达式中的int8不可比较
	// 故会有编译错误。
	// switch 1 + 3 {
	// case value1[0], value1[1]:
	// 	fmt.Println("0 or 1")
	// case value1[2], value1[3]:
	// 	fmt.Println("2 or 3")
	// case value1[4], value1[5], value1[6]:
	// 	fmt.Println("4 or 5 or 6")
	// }

	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// case表达式中的表达式列表，以switch表达式为基准
	// 的自动类型转化，仅在case子表达式的结果值为无类型
	// 常量时才发生。
	switch value2[4] {
	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}

}
