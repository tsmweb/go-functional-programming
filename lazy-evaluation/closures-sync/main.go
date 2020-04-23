// Emulates goroutine-safe lazy evaluation is using closures and the sync-package.
package main

import (
	"fmt"
	"sync"
)

type LazyInt func() int

func main() {
	n := Make(func() int { return 23 }) // Or something more expensive...
	fmt.Println(n())                    // Calculates the 23
	fmt.Println(n() + 42)               // Reuses the calculated value
}

func Make(f func() int) LazyInt {
	var v int
	var once sync.Once
	return func() int {
		once.Do(func() {
			fmt.Println("execute f()")
			v = f()
			f = nil // so that f can now be GC`ed
		})

		return v
	}
}
