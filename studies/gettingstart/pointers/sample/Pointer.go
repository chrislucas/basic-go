/*
	https://go.dev/tour/moretypes/1

	O tipo *T é um ponteiro para um valor de T. Seu valor default é nil.

	O operador & gera um ponteiro para seu operando.
	O operador * denota o valor subjacente ao ponteiro.
*/

package main

import "fmt"

func AnotherTestPointer() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21
	fmt.Println(i) // see the new value of i

	p = &j
	*p = *p / 37
	fmt.Println(j)

	var s *int
	fmt.Println(s)
}

func TestPointer() {
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}

func main() {
	TestPointer()
}
