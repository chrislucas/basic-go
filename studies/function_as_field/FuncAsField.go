// https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

package main

import "fmt"

type Person struct {
	Name  string
	Greet func(message string) string
}

func (p Person) Greeting(message string) string {
	return fmt.Sprintf("message, %s!", p.Name)
}

func main() {
	person := Person{
		Name: "John",
	}

	person.Greet = func(message string) string {
		return fmt.Sprintf("message, %s!", person.Name)
	}

	fmt.Println(person.Greet("Olá"))

	fmt.Println(person.Greeting("Olá"))

}
