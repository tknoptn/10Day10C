// when we want to get fib(9) -- 5 times we calc fib(5) -- but we don't want
// save the value in cache like map in go - we call it a dictionary in python // in java we implement map using hashmap

package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fib(pos int) int {
	if pos == 0 {
		return 0
	}

	if pos == 1 || pos == 2 {
		return 1
	}

	if val, ok := cache[pos]; ok {
		return val
	}

	result := fib(pos-1) + fib(pos-2)
	cache[pos] = result
	return result
}

func main() {
	position := 9
	result := fib(position)
	fmt.Println(result)
}
