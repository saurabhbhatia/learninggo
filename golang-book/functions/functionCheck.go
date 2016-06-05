package main

import "fmt"

// custom function called in main
func main() {
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}

// custom function
func average(xs []float64) float64 {
  if len(xs) == 0 {
    panic("this array is blank")
  } else {
    total := 0.0
    for _, v := range xs {
      total += v
    }
    return (total/float64(len(xs)))
  }
}
