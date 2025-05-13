package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func ReaderIntegers(tranform func(string) []int64) []int64 {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return tranform(text)
}

func ReaderValues[T any](tranform func(string) []T) []T {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return tranform(text)
}

func ReaderValue[T any](tranform func(string) T) (T, error) {
	text, error := reader.ReadString('\n')
	if error != nil && error == io.EOF || text == "" {
		var zero T
		return zero, error
	}
	text = strings.TrimSpace(text)
	return tranform(text), nil
}

func TestReaderValuesI() {
	a := ReaderValues(func(text string) []int64 {
		strValues := strings.Split(text, " ")
		intValues := make([]int64, len(strValues))
		for i, strValue := range strValues {
			parsedValue, _ := strconv.ParseInt(strValue, 10, 64)
			intValues[i] = parsedValue
		}
		return intValues
	})

	fmt.Println(a)

	b := ReaderValues(func(text string) []int64 {
		strValues := strings.Split(text, ";")
		intValues := make([]int64, len(strValues))
		for i, strValue := range strValues {
			parsedValue, _ := strconv.ParseInt(strValue, 10, 64)
			intValues[i] = parsedValue
		}
		return intValues
	})

	fmt.Println(b)
}

func main() {
	TestReaderValuesI()
}

func BuilderString[T any](strs []T, sep string) string {
	var sb strings.Builder
	for _, value := range strs {
		sb.WriteString(fmt.Sprintf("%v%s", value, sep))
	}
	return sb.String()
}

func AnotherRunUntilEOF(inv func(string)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		inv(strings.TrimSpace(text))
	}
}

func RunUntilEOF(inv func(string)) {
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		inv(strings.TrimSpace(text))
	}
}
