package main

import "fmt"

func main() {
  const message string = "Hello world"
  fmt.Println(message)

  // makes program fail because const cannot be modified
  message = "yellow world"
  fmt.Println(message)
}
