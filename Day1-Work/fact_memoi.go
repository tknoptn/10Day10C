package main

import (
	"fmt"
)

var cache = make(map[int]int)

func fact(num int) int {
	if num == 1 {
		return 1
	}
	if val, ok := cache[num]; ok {
		return val
	}
	result := num * fact(num-1)
	cache[num] = result
	return result
}

func main() {
	num := 5
	res := fact(num)
	fmt.Println(res)
}
