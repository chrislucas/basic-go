/*
	https://go.dev/tour/moretypes/15
*/

package main

import "fmt"

func main() {
	testAppendSlice1()
}

func testAppendSlice1() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	s = append(s, 0)
	fmt.Println(s, len(s), cap(s))
}
