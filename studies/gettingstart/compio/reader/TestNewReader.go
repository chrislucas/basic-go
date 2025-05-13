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

	stringValues := strings.Split(input, " ")

	var intValues []int
	for _, strVal := range stringValues {
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Println("Invalid input:", strVal)
			return
		}
		intValues = append(intValues, intVal)
	}

	fmt.Println("Array:", intValues)
}
