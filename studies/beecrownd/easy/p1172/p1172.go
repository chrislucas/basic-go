// https://judge.beecrowd.com/en/problems/view/1172
// https://www.codementor.io/@tucnak/using-golang-for-competitive-programming-h8lhvxzt3

// DONE

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var values [10]int
	for i := 0; i < 10; i++ {
		scanf("%d\n", &values[i])
		if values[i] <= 0 {
			values[i] = 1
		}
	}

	strs := make([]string, len(values))

	for i, val := range values {
		strs[i] = fmt.Sprintf("X[%d] = %d", i, val)
	}

	//printf("%s\n", BuilderString(strs, "\n"))
	//printf("%s", JoinString(strs, "\n"))
	//fmt.Println(strings.Join(strs, "\n"))
	fmt.Print(BuilderString(strs, "\n"))
}

func JoinString(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

func BuilderString(strs []string, sep string) string {
	var sb strings.Builder
	for _, value := range strs {
		sb.WriteString(value + sep)
	}
	return sb.String()
}
