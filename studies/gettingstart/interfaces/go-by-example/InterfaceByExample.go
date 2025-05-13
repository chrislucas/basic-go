// https://gobyexample.com/interfaces
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) Area() float64 {
	return r.width * r.height
}

func (r rect) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(s Shape) {
	fmt.Println(s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func DetectCircle(s Shape) {
	// type assertions https://go.dev/tour/methods/15
	if _, ok := s.(circle); ok {
		fmt.Println("It's a circle!")
	} else {
		fmt.Println("It's not a circle!")
	}
}

func DetectCircleTypeSwitch(s Shape) {
	switch s.(type) {
	case circle:
		fmt.Println("It's a circle!")
	default:
		fmt.Println("It's not a circle!")
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	Measure(r)
	Measure(c)
	DetectCircle(r)
	DetectCircle(c)
	DetectCircleTypeSwitch(c)
}
