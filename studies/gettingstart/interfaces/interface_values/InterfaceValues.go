/*
	https://go.dev/tour/methods/11

	Interface values pode ser pensado como uma tupla de um valor e um tipo concreto:
	(value, type)
*/

package main

import (
	"fmt"
	"math"
)

type Executor interface {
	Run()
}

type Task struct {
	name string
}

func (task *Task) Run() {
	fmt.Println("Running task", task.name)
}

type CustomFloat float64

func (f CustomFloat) Run() {
	fmt.Println("Running CustomFloat", f)
}

func main() {
	/*
		Uma interface value armazena um valor de um tipo concreto adjacente especifico
	*/
	var value Executor = &Task{"main task"}
	Describe(value)

	value = CustomFloat(math.Phi)
	Describe(value)
}

func Describe(e Executor) {
	fmt.Printf("(%v, %T)\n", e, e)
}
