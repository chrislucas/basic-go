package main

import (
	"fmt"
	"strings"
)

func main() {
	values := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	strs := make([]string, len(values))
	for i, val := range values {
		strs[i] = fmt.Sprintf("%d", val)
	}
	fmt.Println(strings.Join(strs, "\n"))
}
