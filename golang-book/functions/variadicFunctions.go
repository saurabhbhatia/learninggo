package main

import "fmt"

// ... is the variadic identifier that says an array of args with a specific datatype can be accepted by the function
func add(args...int) int {
  total := 0
  for _, v := range args {
    total += v
  }
  return total
}

func main() {
  fmt.Println(add(1, 2, 3))
}
