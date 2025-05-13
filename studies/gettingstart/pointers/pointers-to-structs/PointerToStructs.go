// https://go.dev/tour/moretypes/4

/*

	Campos de uma struct podem ser acessados via ponteiro para struct

*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	var p *Vertex = &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println("--------------------------------")

	var q = *p
	q.X = 1e3
	fmt.Println(q)
	fmt.Println(v)
	fmt.Println("--------------------------------")

	var u = p
	u.X = 1e3
	fmt.Println(u)
	fmt.Println(p)
	fmt.Println(v)
	fmt.Println("--------------------------------")
}
