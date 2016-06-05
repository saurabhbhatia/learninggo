package main

import "fmt"

// How do you get the memory address of a variable?
// ans : by adding an astrix and printing it out - x *int points towards the memory location of an integer with value x

//how do you assing a value to the pointer
// *x = newValue
// access this value using &x

// how do you create a new pointer
// x := new(int)

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

  // finding square of a number
  x := 1.5
  square(&x)
  fmt.Println(x)

  // swapping integers
  a := 4
  b := 5
  swapInt(&a, &b)
  fmt.Println(a, b)
}
