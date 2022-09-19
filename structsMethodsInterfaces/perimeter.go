package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(height, width float64) float64 {
	return (height + width) * 2
}

func Area(height, width float64) float64 {
	return height * width
}
