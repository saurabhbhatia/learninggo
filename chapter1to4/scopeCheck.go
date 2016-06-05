package main

import "fmt"

var message string = "Hello world"

func main() {
  fmt.Println(message)
  f()
}

func f() {
  message = "Yellow world"
  fmt.Println(message)
}
