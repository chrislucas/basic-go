/*
	https://go.dev/tour/methods/10

	Nao ha um declaracao explicita de implementacao de uma interface
	nao há a palavra reservada "implements"

*/

package main

import "fmt"

type Executor interface {
	Run()
	Stop()
}

type Task struct {
	name string
}

func (task Task) Run() {
	fmt.Println("Running task", task.name)
}

func (task Task) Stop() {
	fmt.Println("Stopping task", task.name)
}

func main() {
	var task Executor = Task{name: "Main Task"}
	task.Run()

	simpleTask := Task{name: "Simple Task"}
	simpleTask.Run()
}
