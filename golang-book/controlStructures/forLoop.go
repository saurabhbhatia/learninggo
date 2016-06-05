package main

import "fmt"

func main() {
  // for loop
  i := 1
  for i <= 10 {
    fmt.Println(i)
    i = i + 1
  }

  // shorthand for loop
  for i := 11; i <= 20; i++ {
    fmt.Println(i)
  }
}
