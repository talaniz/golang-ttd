package structs

import "math"

type Shape interface {
	Area() float64
}

// Rectangle represents a rect obj
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle reps a shape
type Circle struct {
	Radius float64
}

// Triangle reps a shape
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter takes 2 params and returns a perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}

// Area takes a width/height and returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area takes a radius and returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area takes the base/height and returns area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
