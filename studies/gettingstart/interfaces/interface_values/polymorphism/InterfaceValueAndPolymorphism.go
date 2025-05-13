/*
	pesquisa Interface values golang example

	Pedindo um exemplo para IA do google ela respondeu que a Interface Value
	é um recurso que possibilita alcancar o polimorfismo em GO.

	Elas permite que variavels armazenem valores de tipo concretos que iumplementam
	a mesma interface

	Uma interface value conceitualmente armazena 2 componentes, um concreto e um tipo dinamico
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())

	// Tupla (valor, tipo)
	fmt.Printf("Tuple(Type: %T, Value: %v)\n", s, s)
}

func main() {
	c := Circle{5}
	r := Rectangle{2, 2}

	// interface value para armazenar os tipos concretos
	var s Shape

	s = c
	PrintShapeInfo(s)
	s = r
	PrintShapeInfo(s)
}
