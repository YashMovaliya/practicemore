package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	length float64
}

func (r rect) area() float64 {
	return r.width * r.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c1 := circle{4}
	r1 := rect{3, 4}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
