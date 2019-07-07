package main

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main() {
	// example 1
	// 不能调用不可寻址的值得指针方法
	// New("little pig").SetName("monster")

	// example 2
	// 使用自增++自减--语句是有要求的: 表达式的结果是可以寻址的
	// 字典字面量和字典变量索引表达式是不可寻址的，但是这样的表达式
	// 的结果却可以用于自增自减语句中。
	map[string]int{"the": 0, "world": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
}
