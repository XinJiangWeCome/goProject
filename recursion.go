package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (c Rectangle) area() float64 {
	return c.width * c.height
}
func (r Circle) area() float64 {
	return 3.14 * r.radius * r.radius
}
func caseOne() {
	var s Shape
	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积:%f\n", s.area())
	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
