package main

import (
	"fmt"
)

func main() {
	var m map[string]int

	key := "two"
	elem, ok := m[key]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n", key, elem, ok)

	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n", key)

	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。

	// 注意往空字典中加入键值对会引发panic.

}
