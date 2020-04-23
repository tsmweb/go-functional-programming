package main

import "fmt"

func main() {
	// we are currying the add method to create more variations
	var add10 = add(10)
	var add20 = add(20)
	var add30 = add(30)

	fmt.Println(add10(5)) // 15
	fmt.Println(add20(5)) // 25
	fmt.Println(add30(5)) // 35
}

// this is a higher-order-function that returns a function
func add(x int) func(y int) int {
	// A function is returned here as closure
	// variable x is obtained from the outer scope of this method and memorized in the closure
	return func(y int) int {
		return x + y
	}
}
