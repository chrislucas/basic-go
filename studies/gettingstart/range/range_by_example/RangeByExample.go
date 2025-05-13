// https://gobyexample.com/range

package main

import (
	"fmt"
)

func RangeMap() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	RangeMap()
}
