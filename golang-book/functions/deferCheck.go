package main

import "fmt"

func first()  {
  fmt.Println("see this first")
}

func second()  {
  fmt.Println("see this later")
}

func main() {
  defer second()
  first()
}
