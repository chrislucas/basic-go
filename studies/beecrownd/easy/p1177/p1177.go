// https://judge.beecrowd.com/en/problems/view/1177

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...any)        { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...any)         { fmt.Fscanf(reader, f, a...) }
func format(f string, a ...any) string { return fmt.Sprintf(f, a...) }

func main() {
	defer writer.Flush()
	var length = 1000
	var values = make([]int, length)
	var value int
	scanf("%d", &value)
	values[0] = value
	for i := 0; i < length; i++ {
		values[i] = i % value
	}

	strs := make([]string, len(values))

	for i, val := range values {
		strs[i] = format("N[%d] = %d", i, val)
	}

	//printf("%s", strings.Trim(strings.Join(strs, "\n"), "\n"))
	printf("%s", joinStringsExceptLast(strs, "\n"))
}

func joinStringsExceptLast(strs []string, sep string) string {
	if len(strs) <= 1 {
		return strings.Join(strs, sep)
	}
	return strings.Join(strs[:len(strs)-1], sep)
}
