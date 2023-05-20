package main

import (
	"fmt"
)

// Iteration
func fib(pos int) int {
	a := 0
	b := 1
	c := 0
	for i := 2; i < pos; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
func main() {
	position := 9
	result := fib(position)
	fmt.Println(result)
}
