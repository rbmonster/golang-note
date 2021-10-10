package main

import (
	"fmt"
	"golang-note/src/basic/fib"
)

func init() {
	fmt.Println("initing")
}

func main() {
	var res = fib.Fib(10)
	fmt.Println(res)
}
