// https://go.dev/tour/moretypes/2

package main

import (
	"fmt"
)

type Menu struct {
	text string
}

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Menu{text: "Menu 1"})
}
