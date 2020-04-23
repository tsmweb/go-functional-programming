package main

import "fmt"

type F func(i int) int

func (f F) compose(inner F) F {
	return func(i int) int {
		return f(inner(i))
	}
}

func compose2(a F, b F) F {
	return func(i int) int {
		return a(b(i))
	}
}

func compose(fns ...F) F {
	return func(i int) int {
		var res F
		res = fns[0]

		for i := 1; i < len(fns); i++ {
			res = compose2(res, fns[i])
		}

		return res(i)
	}
}

func main() {
	var f1 F = func(i int) int {
		return i * 2
	}

	var f2 F = func(i int) int {
		return i + 1
	}

	var f3 F = func(i int) int {
		return i + 10
	}

	fc1 := f1.compose(f2).compose(f3)
	fmt.Println(fc1(2)) // 26

	fc2 := compose(f1, f2, f3)
	fmt.Println(fc2(2)) // 26
}
