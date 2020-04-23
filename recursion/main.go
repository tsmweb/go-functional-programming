package main

import "fmt"

func main() {
	fmt.Println(factorialLoop(20))    // 2432902008176640000
	fmt.Println(factorialRec(20))     // 2432902008176640000
	fmt.Println(factorialTailRec(20)) // 2432902008176640000
}

// interative approach
func factorialLoop(num int) int {
	result := 1

	for ; num > 0; num-- {
		result *= num
	}

	return result
}

// recursion (stack overflow)
func factorialRec(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorialRec(num-1)
}

// tail recursion (Go doesn’t optimize this)
func factorialTailRec(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}

	return factorial(accumulator*val, val-1)
}
