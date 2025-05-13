// https://go.dev/tour/flowcontrol/12

package main

import "fmt"

/*
A declaração defer adia a execução de uma função até o final do retorno da função atual,
após a execução de todos os retornos adiados.
*/

func Test1() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func Value() int {
	return 1
}

func TestSumDefered() {
	/*
		defer so espera a execucao da funcao, nao o valor
	*/
	defer func() {
		a := Value()
		fmt.Println(1 + a)
	}()
}

func TestStackingDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	TestStackingDefers()
}
