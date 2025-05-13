package main

import (
	"fmt"
	"reflect"
)

func main() {
	testCreateArrays()
}

func testCreateArrays() {
	intValues := []int{1, 2, 3, 4, 5}
	fmt.Println(intValues)
	stringValues := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(stringValues)

	intValues = append(intValues, 6)
	fmt.Println(intValues)

	/*
		Integer
		Object
		---
		Any
		interfaces{}
	*/
}

func test1() {
	var strs [2]string
	strs[0] = "Hello"
	strs[1] = "World"
	fmt.Println(strs[0], strs[1])
	fmt.Println(strs)

	values := [6]int{1, 2, 3}
	fmt.Println(values)
}

func test2() {
	strs := [2]string{"Hello", "World"}
	fmt.Println(strs)
	fmt.Println(&strs[0])
}

func test3() {
	ints := [2]int{0, 1}
	fmt.Println(ints)
	fmt.Println(&ints[0], reflect.TypeOf(&ints[0]).Size())
	fmt.Println(&ints[1], reflect.TypeOf(&ints[1]).Size())
}

func sizeofTypes() {

	/*
		https://go.dev/tour/basics/11
		Type Size 		(bytes)
		int, uint 		4 or 8 (depending on the architecture)
		int8, uint8     1
		int16, uint16 	2
		int32, uint32   4
		int64, uint64   8
		float32         4
		float64         8
	*/

	var x int
	var s string
	var f float64
	var arr [5]int

	fmt.Println("Size of int (using reflect):", reflect.TypeOf(x).Size(), "bytes")
	fmt.Println("Size of string (using reflect):", reflect.TypeOf(s).Size(), "bytes")
	fmt.Println("Size of float64 (using reflect):", reflect.TypeOf(f).Size(), "bytes")
	fmt.Println("Size of array [5]int (using reflect):", reflect.TypeOf(arr).Size(), "bytes")
}
