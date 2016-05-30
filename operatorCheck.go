package main

import "fmt"

func main() {
  x := "hello"
  y := "world"
  fmt.Println(x == y)
  y = "hello"
  fmt.Println(x == y)
}
