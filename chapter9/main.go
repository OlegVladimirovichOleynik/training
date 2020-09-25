package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
type Shape interface {
	area() float64
	perimeter() float64
}
type MultiShape struct {
	shapes []Shape
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l + w) * 2
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
func main() {
	var c = Circle{x: 0, y: 0, r: 5}
	var r = Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	fmt.Println("Плщадь прямоугольника", r.area(), "Площадь круга", c.area())
	fmt.Println("TotalArea", totalArea(&r, &c))
	var shapes = MultiShape{
		shapes: []Shape{&r, &c, &r},
	}
	fmt.Println("MultiShape", shapes.area())
	fmt.Println("Периметр прямоугольника", r.perimeter(), "Периметр круга", c.perimeter())
	fmt.Println("TotalPerimeter", totalPerimeter(&r, &c))
}
