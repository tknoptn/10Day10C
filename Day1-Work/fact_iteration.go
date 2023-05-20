package main

import (
	"fmt"
)

// 5 = 1 * 2 * 3 * 4 * 5 = 120
// Iteration approach
func factorial(num int) int {
	value := 1
	for i := 1; i <= num; i++ {
		value = value * i
	}
	return value
}
func main() {
	num := 5
	res := factorial(num)
	fmt.Println(res)
}
