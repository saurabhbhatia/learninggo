package main

import "fmt"

func main() {
  // function within a function
  add := func (x, y int) int  {
    return x + y
  }
  fmt.Println(add(5,6))

  // multiple calls of closures
  x := 0
  increment := func () int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())

  // closures
  nextEvenNumber := makeEvenGenerator()
  fmt.Println(nextEvenNumber())
  fmt.Println(nextEvenNumber())
}

func makeEvenGenerator() func() uint  {
  i := uint(0)
  return func() (val uint)  {
    val = i
    i += 2
    return
  }
}
