package main

import "fmt"

func main() {
  // slicing an array
  var x []float64
  y := make([]float64, 3)
  arr := [5] float64{ 1, 2, 3, 4, 5 }
  x = arr[2:4]
  y = arr[1:4]
  z := arr[0:2]
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}
