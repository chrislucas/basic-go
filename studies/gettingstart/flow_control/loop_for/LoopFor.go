// https://go.dev/tour/flowcontrol/1

package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func SimpleSumWithFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func ForConitued() {
	sum := 1
	// o início e o pós são opcionais
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
}

func SimpleSumWithWhile() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}
}

func SimpleInfiniteLoop() {
	var value int
	for {
		Scanf("%d", &value)
		if value < 1 {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func main() {
	SimpleInfiniteLoop()
}
