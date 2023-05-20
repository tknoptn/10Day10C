//5 = 5 * 4!
//4 = 4 * 3!

package main

import (
	"fmt"
)

func fact(num int) int {
	if num == 1 {
		return 1
	}
	return num * fact(num-1)
}
func main() {
	num := 5
	res := fact(num)
	fmt.Println(res)
}
