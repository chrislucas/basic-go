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
	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	// "1\n"
	b, _ := reader.ReadString('\n')
	b = strings.TrimSpace(b)

	sa, ea := toInt(a)
	if ea != nil {
		return
	}

	sb, eb := toInt(b)
	if eb != nil {
		return
	}

	fmt.Println("Saída:", (sa + sb))
}

func toInt(value string) (int, error) {
	return strconv.Atoi(value)
}
