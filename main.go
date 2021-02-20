package main

import "fmt"

// initialise shape of type interface and call getArea()
type shape interface {
	getArea() float64
}

//extends struct data type and manipulate to add property of sideLength
type square struct {
	sideLength float64
}

//extends struct data type and maipulate to add properties; base & height
type triangle struct {
	height float64
	base   float64
}

func main() {

	//Declare variables and assign measurements
	square := square{12}
	triangle := triangle{4, 10}

	//Pass variables into printArea function
	printArea(square)
	printArea(triangle)
}

//Function that takes a shape in as argument that satisfies either of the getArea function.
func printArea(s shape) {
	fmt.Println("The area of shape is", s.getArea())
}

//function with receiver of square of type struct that returns area of type float64
func (s square) getArea() float64 {
	area := s.sideLength * 2
	return area
}

//Function with receiver of triangle of type struct that returns the area of type float64
func (t triangle) getArea() float64 {
	area := (t.base / 2) * t.height
	return area
}
