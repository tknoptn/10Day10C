//** now with recursion
//fib(9) = fib(7)+fib(8)
//fib(7) = fib(6)+fib(5)

package main

import (
	"fmt"
)

func fib(pos int) int {
	if pos == 0 {
		return 0
	}
	if pos == 1 || pos == 2 {
		return 1
	}
	return fib(pos-1) + fib(pos-2)
}

func main() {
	position := 9
	result := fib(position)
	fmt.Println(result)
}
