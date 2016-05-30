package main

import "fmt"

func main() {
  // exercise 1
  var message string = "Hello world"
  secondMessage := "Yellow world"

  fmt.Println(message)
  fmt.Println(secondMessage)

  // exercise 2
  x := 5;x += 1;
  fmt.Println(x)

  // exercise 3
  fmt.Println("Enter a temperature in farenheit:")
  var input float64
  fmt.Scanf("%f", &input)
  var output float64 = (((input * 9) / 5) + 32)
  fmt.Println(output, "degree F")

  // exercise 4
  fmt.Println("Convert a feet value into metres:")
  var metresInput float64
  fmt.Scanf("%f", &metresInput)
  var metresOutput float64 = (metresInput * .348)
  fmt.Println(metresOutput, "metres")
}
