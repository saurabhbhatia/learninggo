package main

import "fmt"

func square(x *float64) {
  *x = *x * *x
}

// swap 2 integers
func swapInt(a *int, b *int) {
  c := *a
  d := *b
  *a = d
  *b = c
}

func main() {
  x := 1.5
  square(&x)
  fmt.Println(x)
  a := 4
  b := 5
  swapInt(&a, &b)
  fmt.Println(a, b)
}
