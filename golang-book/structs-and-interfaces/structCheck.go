package main

import "fmt"
import "math"

// set type circle
type Circle struct {
  x, y, r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

// set type rectangle
type Rectangle struct {
  x1, y1, x2, y2 float64
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

func main() {
  // var c Circle
  // c := Circle{x: 0, y: 0, r: 5}
  c := Circle{0,0,5}
  // fields
  fmt.Println("fields", c.x, c.y, c.r)
  fmt.Println("circle area", (c.area()))

  r := Rectangle{0,0,10,20}
  fmt.Println("rectangle area", (r.area()))
}
