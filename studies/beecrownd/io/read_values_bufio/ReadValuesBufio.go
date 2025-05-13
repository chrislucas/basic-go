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

	values := strings.Split(input, " ")
	array := make([]int, len(values))

	for i, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			return
		}
		array[i] = num
	}

	fmt.Println("Array:", array)
}
