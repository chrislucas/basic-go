package main

/*
	Slices
	https://go.dev/tour/moretypes/7

	Arrays tem tamanho fixo e slices tem tamanhos dinamicos.

	O Tipo [] T é um slice com elementos do tipo T.
	Um slice é formado ao especificar 2 indices, low e high bound
		- a[low: high]
		- conjunto semiaberto, inclui o primeiro elemento, mas exclui o ultimo
		- a[1:4] => [1, 2, 3]
		- a[1:] => [1, 2, 3, 4]
		- a[:3] => [0, 1, 2, 3]

	Make Slices
	https://go.dev/tour/moretypes/13
*/

import "fmt"

func main() {
	makeSlices()
}

func makeSlices() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	fmt.Println(a[1:4]) // do indice 1 ate o indice 4 exclusivo
	fmt.Println(a[:3])  // ate o indice 3 exclusivo
	fmt.Println(a[4:])  // do indice 4 ate o final
}
