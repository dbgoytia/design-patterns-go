package main

import (
	"fmt"
	"math"
)

// Our interface is for geometrical objects
type geometry interface {
	area() float64
	perim() float64
}

// We will implement this interface on the rectangle and circle types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface in Go, you must implement all
// methods of the interface. For example for rects:
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// For circles:
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type, we can call
// methods that are in the named interface. For example,
// Here we have a generig "measure" function that will work
// on any "geometry".

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// We can implement geometry
func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
