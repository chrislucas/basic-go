package main

import "fmt"

func binsearch(array []int, value int) int {
	lf := 0
	ri := len(array) - 1
	for lf <= ri {
		mi := (ri-lf)/2 + lf
		fmt.Println(fmt.Sprintf("%d, %d, %d", lf, ri, mi))
		if array[mi] == value {
			return mi
		} else if array[mi] > value {
			ri = mi - 1
		} else {
			lf = mi + 1
		}
	}

	return -1
}

func testBinSearch() {
	array := [][]int{
		{1, 2, 3, 4, 4, 5, 10, 20, 200, 312, 500},
	}
	fmt.Println(binsearch(array[0], 20))
	fmt.Println(binsearch(array[0], 500))
	fmt.Println(binsearch(array[0], 1))
	fmt.Println(binsearch(array[0], 501))
	fmt.Println(binsearch(array[0], 0))
}

func main() {
	testBinSearch()
}
