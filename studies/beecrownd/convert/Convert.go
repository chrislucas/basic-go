// https://pkg.go.dev/strconv

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseInt("123", 10, 64))
	fmt.Println(strconv.FormatInt(10, 16))
	fmt.Println(strconv.FormatInt(-10, 16))
	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println(strconv.FormatInt(-10, 2))
}
