// https://go.dev/tour/methods/9
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var absolute Abser
	c := CustomFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	absolute = c  // a CustomFloat implements Abser
	absolute = &v // a *Vertext implements Abser
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// absolute = v

	fmt.Println(absolute)
	fmt.Println(absolute.Abs())
}

type CustomFloat float64

func (f CustomFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

// *Vertext implements Abser
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
