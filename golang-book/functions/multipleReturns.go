package main

import "fmt"

// return multiple values
func main() {
  x, y := f()
  fmt.Println(x, y)
}

func f() (int, int)  {
  return f1(), f2()
}

func f1() int {
  return 5
}

func f2() int {
  return 6
}
