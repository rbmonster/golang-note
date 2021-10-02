package main

import (
	"fmt"
	"golang-note/src/basic"
)

func init() {
	fmt.Println("initing")
}

func main() {
	var res = basic.Fib(10)
	fmt.Println(res)
}
