package main

import "fmt"

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}

func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[1] = []string{"o", "p", "q"}
	return a
}

func main() {
	// example 1, 数组为值类型，传给函数时进行值拷贝
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2) // [a x c]
	fmt.Printf("The original array: %v\n", array2) // [a b c]
	fmt.Println()

	// exmaple 2, 切片为引用类型，传给函数是做潜拷贝，不会拷贝底层数组
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice2: %v\n", slice2) // [x i z]
	fmt.Printf("The original slice1: %v\n", slice1) // [x i z]
	fmt.Println()

	// exmaple 3
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	// [[d e f] [o p q] [j k  l]]
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	// [[d e f] [g s i] [j k l]]
	fmt.Printf("The original complex array: %v\n", complexArray1)
}
