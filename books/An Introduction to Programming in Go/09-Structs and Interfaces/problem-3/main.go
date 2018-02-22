package main

import "fmt"

type Shape interface {
	perimeter() float64
}
type MultiShape struct {
	shapes []Shape
}

func totalPerimeter(m *MultiShape) (perimeter float64) {
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return
}

type Circle struct {
	r float64
}
type Rectangle struct {
	w float64
	l float64
}

func (c *Circle) perimeter() float64 {
	return 2 * 3.14 * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.w + r.l)
}

func main() {
	c := Circle{10}
	r := Rectangle{10, 20}
	shapes := MultiShape{[]Shape{&r, &c}}
	fmt.Println(totalPerimeter(&shapes))
}
