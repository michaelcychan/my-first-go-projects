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

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return (c.Radius * 2 * math.Pi)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (tri Triangle) Area() float64 {
	return tri.Base * tri.Height / 2
}
