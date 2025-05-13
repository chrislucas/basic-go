// https://gobyexample.com/pointers

package main

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	println("initial:", i)

	zeroval(i)
	println("zeroval:", i)

	// zeroptr muda o valor de i pois passamos o endereço de memória, um ponteiro para a variavel i
	zeroptr(&i)
	println("zeroptr:", i)

	// & passa o endereço de memória, dando um ponteiro para a variavel i
	println("pointer:", &i)
}
