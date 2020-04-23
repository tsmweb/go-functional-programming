package main

import "fmt"

func main() {
	/*
		executing add
		executing multiply
		8
		executing add
		executing multiply
		16
	*/
	fmt.Println(addOrMultiply(true, add(4), multiply(4)))  // 8
	fmt.Println(addOrMultiply(false, add(4), multiply(4))) // 16

	/*
		executing add
		8
		executing multiply
		16
	*/
	fmt.Println(addOrMultiplyLazy(true, add, multiply, 4))  // 8
	fmt.Println(addOrMultiplyLazy(false, add, multiply, 4)) // 16
}

func add(x int) int {
	fmt.Println("executing add") // this is printed since the functions are evaluated first
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply") // this is printed since the functions are evaluated first
	return x * x
}

func addOrMultiply(add bool, onAdd, onMultiply int) int {
	if add {
		return onAdd
	}

	return onMultiply
}

// This is now a higher-order-function hence evaluation of the functions are delayed in if-else
func addOrMultiplyLazy(add bool, onAdd, onMultiply func(t int) int, t int) int {
	if add {
		return onAdd(t)
	}

	return onMultiply(t)
}
