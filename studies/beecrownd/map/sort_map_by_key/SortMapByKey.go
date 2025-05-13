package main

import (
	"fmt"
	"sort"
)

func SortMapByKey[K comparable, V any](data map[K]V, lessThan func(a, b K) bool) []K {

	keys := make([]K, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return lessThan(keys[i], keys[j])
	})

	return keys
}

func SortingMapByKey[K comparable, V any](data map[K]V, lessThan func(a, b K) bool) map[K]V {

	keys := make([]K, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return lessThan(keys[i], keys[j])
	})

	new_map := make(map[K]V)
	for _, key := range keys {
		new_map[key] = data[key]
	}

	return new_map
}

func SortMapByValue[K comparable, V comparable](data map[K]V, lessThan func(a, b K) bool) map[K]V {

	return data
}

func TestMapString(data map[string]int) {

	keys := SortMapByKey(data, func(a, b string) bool {
		return a > b
	})

	for _, key := range keys {
		fmt.Println(key, data[key])
	}
}

func TestMapInt(data map[int]string) {
	keys := SortMapByKey(data, func(a, b int) bool {
		return a > b
	})
	for _, key := range keys {
		fmt.Println(key, data[key])
	}
}

func TestMapFloat(data map[float64]string) {
	keys := SortMapByKey(data, func(a, b float64) bool {
		return a < b
	})

	//new_map := make(map[float64]string)
	for _, key := range keys {
		//new_map[key] = data[key]
		fmt.Println(key, data[key])
	}

	fmt.Println("----")
	fmt.Println(SortingMapByKey(data, func(a, b float64) bool {
		return a < b
	}))
}

func main() {
	TestMapString(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	})

	TestMapInt(map[int]string{
		1: "one",
		4: "four",
		3: "three",
		2: "two",
		5: "five",
	})

	TestMapFloat(map[float64]string{
		5.1: "five",
		1.0: "one",
		4.4: "four",
		3.7: "three",
		2.3: "two",
	})
}
