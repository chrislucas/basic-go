package main

import (
	"fmt"
	"sort"
)

func InPlaceSortingMapByKey[K comparable, V any](data *map[K]V, lessThan func(a, b K) bool) {

	keys := make([]K, 0, len(*data))
	for key := range *data {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return lessThan(keys[i], keys[j])
	})

	new_map := make(map[K]V)
	for _, key := range keys {
		new_map[key] = (*data)[key]
	}

	*data = new_map
}

func main() {

	m := map[int]string{
		1: "one",
		4: "four",
		3: "three",
		2: "two",
		5: "five",
	}

	fmt.Println(m)
	InPlaceSortingMapByKey(&m, func(a, b int) bool {
		return a > b
	})

	fmt.Println(m)
}
