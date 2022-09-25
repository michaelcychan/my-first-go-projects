package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (c Circle) Perimeter() float64 {
	return (c.Radius * 2 * math.Pi)
}
