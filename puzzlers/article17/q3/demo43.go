package main

import "fmt"

func main() {
	// 示例1。case子句中有重复的值
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	//switch value3[4] { // 这条语句无法编译通过。
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	// example2: 绕过了编译，不会应为case字句值重复而报编译错误
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}

	// example3:
	// value6 := interface{}(byte(127))
	// 类型检测，byte与uint8重复
	// switch t := value6.(type) {
	// case uint8, uint16:
	// 	fmt.Println("uint8 or uint16")
	// case byte:
	// 	fmt.Println("byte")
	// default:
	// 	fmt.Println("unsupported type: %T", t)
	// }
}
