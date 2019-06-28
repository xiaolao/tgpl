package main

import "fmt"

func main() {

	var srcInt = int16(-255)

	fmt.Printf("The complement of srcInt: %b (%b)\n", uint16(srcInt), srcInt)

	dstInt := int8(srcInt)
	fmt.Printf("The complement of dstInt: %b (%b)\n", uint8(dstInt), dstInt)
	fmt.Printf("The value of dstInt: %d\n", dstInt)
	fmt.Println()

	fmt.Printf("The Replacement Character: %s\n", string(-1))
	fmt.Printf("The Union codepoint of Replacement Character: %U\n", '�')
	fmt.Println()

	srcStr := "你好"
	fmt.Printf("The string: %q\n", srcStr)
	fmt.Printf("The hex of %q: %x\n", srcStr, srcStr)
	fmt.Printf("The byte slice od %q: %x\n", srcStr, []byte(srcStr))
}
