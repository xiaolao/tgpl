package main

import "fmt"

// solution 1
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type operateFunc func(x int, y int) int

type calculateFunc func(x int, y int) (int, error)

// solution 2
func genCalculator(op operateFunc) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// solution 1: pass func as in args.
	x, y := 12, 23
	op := func(x, y int) {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	// solution 2: by generate caculate function.
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
