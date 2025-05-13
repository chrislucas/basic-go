// https://blog.logrocket.com/comprehensive-guide-data-structures-go/

package main

import "fmt"

func MakeArray() {
	values := make([]int, 5)
	values[0] = 1
	values[1] = 2
	values[2] = 3
	values[3] = 4
	values[4] = 5
	fmt.Println(values)
}

func CreateArray() {
	values := []int{}
	fmt.Println(len(values))
	values = append(values, 1)
	values = append(values, 2)
	values = append(values, 3)
	fmt.Println(len(values))
}

func main() {
	CreateArray()
}
