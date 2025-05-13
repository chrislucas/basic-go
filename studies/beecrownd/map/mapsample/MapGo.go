// https://blog.logrocket.com/comprehensive-guide-data-structures-go/

package main

import (
	"fmt"
)

func CreateArray() {
	var intValues1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intValues1)

	var intValues2 [10]int
	intValues2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intValues2)
}

func CreateAndAssingingValues() {
	values := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(values)
}

func CreateNestedArray() {
	var nestedArray = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nestedArray)
	// fmt.Println(nestedArray[0][0])

	for i, v := range nestedArray {
		fmt.Println(i, v)
	}

	for _, v := range nestedArray {
		for j, w := range v {
			if j == 0 {
				fmt.Print(w)
			} else {
				fmt.Print(" ", w)
			}
		}
		fmt.Println()
	}
}

// maps https://blog.logrocket.com/comprehensive-guide-data-structures-go/

func CreateMap() {
	// https://www.dotnetperls.com/map-string-slice-go
	var mapValues = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Println(mapValues)
}

func InitmapWithMakeFunc() {
	studentAge := make(map[string]int)
	studentAge["john"] = 32
	studentAge["bob"] = 31
	studentAge["alice"] = 28
	studentAge["mary"] = 29
	studentAge["jane"] = 35
	fmt.Println(studentAge)
}

func InitMapOfSlices() {
	values := make(map[string][]int)
	values["one"] = []int{1}
	values["two"] = []int{2, 3}
	values["three"] = []int{4, 5, 6}
	values["four"] = []int{7, 8, 9, 10}
	values["five"] = []int{11, 12, 13, 14, 15}
	fmt.Println(values)

	values["five"] = append(values["five"], 16)
	fmt.Println(values)
}

func InitMap() {
	m := map[string][]int{
		"one":   {1},
		"two":   {2, 3},
		"three": {4, 5, 6},
		"four":  {7, 8, 9, 10},
		"five":  {11, 12, 13, 14, 15},
	}

	fmt.Println(m)
}

func main() {
	InitMap()
}
