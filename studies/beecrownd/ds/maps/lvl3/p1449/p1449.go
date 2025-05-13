// https://judge.beecrowd.com/en/problems/view/1449

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {

}

func BuilderString[T any](strs []T, sep string) string {
	var sb strings.Builder
	for _, value := range strs {
		sb.WriteString(fmt.Sprintf("%v%s", value, sep))
	}
	return sb.String()
}

func Stringify[T any](format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
