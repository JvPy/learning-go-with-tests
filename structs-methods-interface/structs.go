package structs

import "math"

//Interfaces

//Shape - An interface for calculating geometrical methods
type Shape interface {
	Area() float64
}

//Structs

//Rectangle - Struct for defining a Rectangle obj
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle - Struct for defining a Circle obj
type Circle struct {
	Radius float64
}

//Triangle - Struct for defining a Triangle obj
type Triangle struct {
	Base   float64
	Height float64
}

//Functions

//Perimeter - Calculate the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area - Calculate the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
