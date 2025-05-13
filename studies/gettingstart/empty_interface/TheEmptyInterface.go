package main

import("fmt")

/*
	Empty interface
	https://go.dev/tour/methods/14
*/

func main() {
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}