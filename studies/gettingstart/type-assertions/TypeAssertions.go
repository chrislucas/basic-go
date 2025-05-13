// https://go.dev/tour/methods/15

// Switch https://go.dev/tour/methods/16
package main

import "fmt"

func TypeSwitch(i any) {
	switch i.(type) {
	case int:
		fmt.Println("It's an int!")
	case string:
		fmt.Println("It's a string!")
	default:
		fmt.Printf("Type %T!\n", i)
	}
}

func main() {
	var i any = 42

	var _, ok = i.(string)
	fmt.Println(ok)
	fmt.Println(i.(int))

	TypeSwitch(42)
	TypeSwitch("hello")
	TypeSwitch(true)
	TypeSwitch(3.14)
	TypeSwitch(complex(1, 2))
	TypeSwitch([]int{1, 2, 3})
	TypeSwitch(map[string]int{"one": 1})
	TypeSwitch(make(chan int))
	TypeSwitch(struct{}{})
	TypeSwitch(nil)
}
