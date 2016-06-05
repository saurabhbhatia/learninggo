package main

import "fmt"

func main() {
  // array basics
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  // array type casting
  var y [5] float64
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83

  var total float64 = 0
  for i := 0; i < 5; i++ {
    total += y[i]
  }
  fmt.Println(total / float64(len(y)))

  // un-used variable definition
  var arrayTotal float64 = 0
  for _, value := range y {
    arrayTotal += float64(value)
  }
  fmt.Println(total / float64(len(y)))

  // array declaration shorthand
  z := [5]float64{ 98, 93, 77, 82, 83 }
  fmt.Println(z)
}
