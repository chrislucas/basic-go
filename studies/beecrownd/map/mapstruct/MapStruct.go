package main

import "fmt"

type Menu struct {
	text string
}

func MenuMap() {

	/*
		map[k][v]
		1 -> 1.1
		  -> 1.2
		  -> 1.3
		2 -> 2.1
		  -> 2.2
		  -> 2.3
		//

		map[string][]int
		"abc" -> [1, 2, 3]
		"def" -> [4, 5, 6]
		"ghi" -> [7, 8, 9]
		"jkl" -> [10, 11, 12]
		"mno" -> [13, 14, 15]


		A -> B, C, D
		B -> C, D, F
		C -> Z, F

		Lista de adjacência

		map[Pessoa][]Pessoa

		hash O(1)
		  0 {1, 2, 3}
		  1 {1, 2. 3}

		  "abc"
		  UUID -> 123
		  		  456
	*/

	// [Menu] -> []Menu
	var menu2 = Menu{text: "Menu 2"}

	var menus = map[Menu][]Menu{
		{text: "Menu 1"}: {
			Menu{text: "Menu 1.1"},
			Menu{text: "Menu 1.2"},
		},
		menu2: {
			Menu{text: "Menu 2.1"},
			Menu{text: "Menu 2.2"},
		},
		{text: "Menu 3"}: {
			Menu{text: "Menu 3.1"},
			Menu{text: "Menu 3.2"},
		},
	}

	//fmt.Println(menus)
	fmt.Println(menus[Menu{text: "Menu 2"}])
	fmt.Println(menus[menu2])

	fmt.Printf("%t\n", &Menu{text: "Menu 2"} == &menu2)
	fmt.Println(menus[Menu{text: "Menu 4"}])
}

func InitMenu() {
	menu := Menu{
		text: "Menu 1",
	}
	fmt.Print(menu)
}

func OrderMap() {
	m := map[string][]int{
		"one":   {1},
		"two":   {2, 3},
		"three": {4, 5, 6},
		"four":  {7, 8, 9, 10},
	}
	fmt.Println(m)

	ints := map[int]string{
		1: "one",
		4: "four",
		3: "three",
		2: "two",
	}
	fmt.Println(ints)
}

func main() {
	OrderMap()
}
