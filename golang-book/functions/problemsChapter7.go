package main

import "fmt"

func sum(args...float64) float64  {
  total := 0.0
  for _, v := range args {
    total += v
  }
  return total
}

// x... is used to unpack an array into arguments
func main() {
  arr := [5]float64{5,6,7,8,9}
  x := arr[2:5]
  fmt.Println(sum(x...))
}
