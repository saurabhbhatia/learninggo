package main

import "fmt"

func main() {
  for i := 1; i <= 10; i++ {
    if i % 2 == 0 {
      fmt.Println(i," is even")
    } else {
      fmt.Println(i," is odd")
    }
  }

  for i := 1; i <= 20; i++ {
    if i % 2 == 0 {
      fmt.Println(i, " is divisible by 2")
    } else if i % 3 == 0 {
      fmt.Println(i, " is divisible by 3")
    } else if i % 4 == 0 {
      fmt.Println(i, " is divisible by 4")
    }
  }
}
