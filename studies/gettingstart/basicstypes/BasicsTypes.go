package main

import (
	"fmt"
	"math/cmplx"
)

var (
	booleanVar bool       = false
	maxInt     uint64     = 1<<64 - 1
	z          complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", booleanVar, booleanVar)
	fmt.Printf("Type: %T, Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
