package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

// https://judge.beecrowd.com/en/problems/view/1070
// DONE
func main() {
	var value int
	Scanf("%d", &value)
	SolutionB(value)
}

func SolutionB(value int) {
	limit := 6
	values := make([]int, limit)

	if value&1 == 1 {
		values[0] = value
	} else {
		value++
	}

	values[0] = value
	for i := 1; i < 6; i++ {
		value += 2
		values[i] = value
	}

	fmt.Print(BuilderString(values, "\n"))
}

func SolutionA(value int) {
	limit := 5
	if value&1 == 1 {
		fmt.Println(value)
		value += 2
		limit--
	} else {
		value++
	}
	fmt.Println(value)
	for i := 0; i < limit; i++ {
		value += 2
		fmt.Println(value)
	}
}

func BuilderString[T any](strs []T, sep string) string {
	var sb strings.Builder
	for _, value := range strs {
		sb.WriteString(fmt.Sprintf("%v%s", value, sep))
	}
	return sb.String()
}

func Stringify[T any](value T) string {
	return fmt.Sprintf("%v", value)
}
