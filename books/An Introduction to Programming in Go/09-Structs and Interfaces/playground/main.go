package main

// func distance(x1, y1, x2, y2 float64) float64 {
// 	a := x2 - x1
// 	b := y2 - y1
// 	return math.Sqrt(a*a + b*b)
// }
// func rectangleArea(x1, y1, x2, y2 float64) float64 {
// 	l := distance(x1, y1, x1, y2)
// 	w := distance(x1, y1, x2, y1)
// 	return l * w
// }
// func circleArea(x, y, r float64) float64 {
// 	return math.Pi * r * r
// }

// func main() {
// 	var rx1, ry1 float64 = 0, 0
// 	var rx2, ry2 float64 = 10, 10
// 	var cx, cy, cr float64 = 0, 0, 5
// 	fmt.Println(rectangleArea(rx1, ry1, rx2,
// 		ry2))
// 	fmt.Println(circleArea(cx, cy, cr))
// }
// type Circle struct {
// 	x, y, r float64
// }

// func main() {

// 	// var c Circle
// 	// c.x = 1.2
// 	// c.y = 1.2
// 	// c.r = 2
// 	// c := new(Circle) // pointer
// 	// c := Circle{x: 1.2, y: 1.2, r: 2}
// 	c := Circle{1.2, 1.2, 2}

// 	fmt.Println(c.area())
// }
// func (c *Circle) area() float64 {
// 	return math.Pi * c.r * c.r
// }

// func circleArea(c Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// type Rectangle struct {
// 	x1, y1, x2, y2 float64
// }

// func (r *Rectangle) area() float64 {
// 	l := distance(r.x1, r.y1, r.x1, r.y2)
// 	w := distance(r.x1, r.y1, r.x2, r.y1)
// 	return l * w
// }
// func distance(x1, y1, x2, y2 float64) float64 {
// 	a := x2 - x1
// 	b := y2 - y1
// 	return math.Sqrt(a*a + b*b)
// }

// func main() {
// 	sq := Rectangle{1, 1, 10, 10}
// 	fmt.Println(sq.area())

// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Talk() {
// 	fmt.Println("Hi, my name is", p.Name)
// }

// type Android struct {
// 	Person
// 	Model string
// }

// func main() {
// 	// p := Person{Name: "Taey"}
// 	// android := Android{p, "OK101"}
// 	// android.Talk()
// 	a := new(Android)
// 	a.Talk()
// }

type Shape interface {
	area() float64
}
type MultiShape struct {
	shapes []Shape
}

func totalArea(m *MultiShape) float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
