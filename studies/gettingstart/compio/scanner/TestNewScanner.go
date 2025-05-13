package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read a line of space-separated integers
	scanner.Scan()
	line := scanner.Text()
	stringValues := strings.Split(line, " ")

	intValues := make([]int, len(stringValues))
	for i, strVal := range stringValues {
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			return
		}
		intValues[i] = intVal
	}

	fmt.Println("Integers:", intValues)

	// Read a line of space-separated strings
	scanner.Scan()
	line = scanner.Text()
	stringValues = strings.Split(line, " ")

	fmt.Println("Strings:", stringValues)
}
