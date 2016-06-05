package main

import "fmt"

func main() {
  var x [5]int
  for i := 0; i < 5; i++ {
    x[i] = i + 1
  }
  // from array
  fmt.Println("fourth element of the array is", x[3])
  // slice
  fmt.Println("fourth element of array using slice is", x[3:4])

  // check the length of the slice
  y := make([]int, 3, 9)
  fmt.Println(len(y))

  // what's the output
  z := [6]string{"a","b","c","d","e","f"}
  fmt.Println(z[2:5])

  // array looping exercise - find the smallest number in the array
  a := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }
  var number int = a[0]
  for j := 0; j < len(a); j++ {
    if a[j] < number{
      number = a[j]
    }
  }
  fmt.Println("the smallest number in the arry is", number)
}
