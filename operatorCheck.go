package main

import "fmt"

func main() {
  // auto-assign the variable datatype
  x := "hello"
  y := "world"
  fmt.Println(x == y)
  y = "hello"
  fmt.Println(x == y)
}
