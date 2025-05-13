package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"a", "b", "c"}
	t := strings.Join(s, "\n")
	fmt.Println(t)
}
