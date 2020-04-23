package main

import "fmt"

var holder = map[string]int{}

func main() {
	pf := sum(2, 3)
	fmt.Println(pf) // 5

	se := sum2(2, 3)
	fmt.Println(se)     // 5
	fmt.Println(holder) // map[2+3:5]
}

// pure function
func sum(a, b int) int {
	return a + b
}

// side effect that affects an external state
func sum2(a, b int) int {
	c := a + b
	holder[fmt.Sprintf("%d+%d", a, b)] = c
	return c
}
