package main

import "fmt"

func main() {
  // append slices
  sliceX := []int{ 1, 2, 3 }
  sliceY := append(sliceX, 4, 5)
  fmt.Println(sliceX, sliceY)

  // part 2
  sliceA := []int{1, 2, 3}
  sliceB := make([]int, 2)
  copy(sliceA, sliceB)
  fmt.Println(sliceA, sliceB)
}
