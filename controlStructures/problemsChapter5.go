package main

import "fmt"

func main() {

  // output small
  i := 10
  if i > 10 {
    fmt.Println("Big")
  } else {
    fmt.Println("Small")
  }

  // divisible by 3
  for i := 1; i <= 100; i ++ {
    if i % 3 == 0 {
      fmt.Println(i, "is divisible by 3")
    }
  }

  // 3 & 5 fizzbuzz
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 != 0 {
      fmt.Println("fizz")
    } else if i % 3 != 0 && i % 5 == 0 {
      fmt.Println("buzz")
    } else if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("fizzbuzz")
    }
  }
}
