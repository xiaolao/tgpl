package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// example 1:
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name()) // little pig
	// 接口类型本身是不可以初始化
	// 接口被赋值会将数据保存在数据结构iface中
	// iface有两个指针一个指向动态类型，一个指向动态值
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name()) // monster
	fmt.Printf("This pet is a %s, the name is %q. \n",
		pet.Category(), pet.Name()) // dog, little dog
	fmt.Println()

	// example 2:
	dog1 := Dog{"little pig"}
	fmt.Printf("the name of the first dog is %q.\n", dog1.Name())  // little pig
	dog2 := dog1                                                   // 将一个变量赋值给另一个变量,会copy副本
	fmt.Printf("the name of the second dog is %q.\n", dog2.Name()) // little pig
	dog1.name = "monster"
	fmt.Printf("the name of the first dog is %q.\n", dog1.Name())  // monster
	fmt.Printf("the name of the second dog is %q.\n", dog2.Name()) // little pig
	fmt.Println()

	// example 3:
	dog = Dog{"little pig"}
	fmt.Printf("The dog name is %q.\n", dog.Name()) // little pig
	pet = &dog                                      // pet 将&dog的值copy到iface数据接口中
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name()) // monster
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name()) // dog, monster &dog的副本name指针对应的value还是dog的
}
