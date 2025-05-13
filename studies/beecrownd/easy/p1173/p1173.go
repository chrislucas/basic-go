// https://judge.beecrowd.com/en/problems/view/1173
// DONE

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
// func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

func main() {
	//defer writer.Flush()
	var values [10]int
	var value int
	Scanf("%d", &value)
	values[0] = value
	for i := 1; i < 10; i++ {
		values[i] = values[i-1] * 2
	}

	strs := make([]string, len(values))

	for i, val := range values {
		strs[i] = fmt.Sprintf("N[%d] = %d", i, val)
	}

	fmt.Print(BuilderString(strs, "\n"))
}

func BuilderString(strs []string, sep string) string {
	var sb strings.Builder
	for _, value := range strs {
		sb.WriteString(fmt.Sprintf("%s%s", value, sep))
	}
	return sb.String()
}
