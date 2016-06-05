package main

import "fmt"

func zero(valPtr *int) {
  *valPtr = 0
}

func one(xPtr *string) {
  *xPtr = "Got me"
}

func main() {
  val := 5
  zero(&val)
  fmt.Println(val)

  xPtr := new(string)
  one(xPtr)
  fmt.Println(*xPtr)
}
