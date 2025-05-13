package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	stringValues := strings.Split(input, ", ")
	var intValues []int
	for _, value := range stringValues {
		intValue, err := toInt(value)
		if err != nil {
			return
		}
		intValues = append(intValues, intValue)
	}
	fmt.Println("Array:", intValues)
}

func toInt(value string) (int, error) {
	return strconv.Atoi(value)
}
