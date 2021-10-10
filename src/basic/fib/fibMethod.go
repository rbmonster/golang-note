package fib

import "fmt"

func init() {
	fmt.Println("basic initing")
}

func Fib(n int) int {
	var x, y = 0, 1
	for i := 2; i <= n; i++ {
		x, y = y, x+y
	}
	return y
}
