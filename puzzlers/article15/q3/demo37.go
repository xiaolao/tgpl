package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// example 1
	// 指针类型可以和unsafe.Pointer相互转化
	// unsafe.Pointer可以和uintptr相互转化
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))
	fmt.Println(dogPtr) // 获取结构体dog的内存地址

	// unsafe.OffsetOf: 该函数返回结构体与结构体中指定字段
	// 的字节距离，传入参数必须是structValue.name的格式。
	namePtr := dogPtr + unsafe.Offsetof(dogP.name) // 获取字段name的内存地址
	fmt.Println(namePtr)
	fmt.Printf("The Offset of dogP and nameP is %v.\n", unsafe.Offsetof(dogP.name))

	// *string为string的指针类型
	// 此处将uintptr类型转化成unsafe.pointer类型再转化成*string类型
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(dogP.name)? %v\n", nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster" // 获取name字段指针后，可直接通过nameP来修改指针对应的值
	fmt.Printf("The name of dog is %q.\n", *nameP)
	fmt.Println()

	// example 2: 将uintptr转成unsafe.Pointer,再转成*int指针
	// 不同数据类型在内存中都是以二进制的方式表示的，由指针指向
	// 不同类型指针对应不同的读取方式，返回相应类型的数据值。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)
}
