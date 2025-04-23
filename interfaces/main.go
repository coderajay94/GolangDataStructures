package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	length, bredth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.bredth * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	fmt.Println("interfaces")

	rectangle1 := Rectangle{length: 12.3, bredth: 13.3}
	circle1 := Circle{radius: 12.22}

	fmt.Println(CalculateArea(rectangle1))
	fmt.Println(CalculateArea(circle1))

	
}
