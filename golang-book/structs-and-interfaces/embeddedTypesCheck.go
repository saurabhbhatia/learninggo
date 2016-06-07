package main

import "fmt"

type Person struct {
  Name string
}

type Android struct {
  Person
  Model string
}

func (p *Person) Talk() {
  p.Name = "Andrew"
  fmt.Println("My name is", p.Name)
}

func main() {
  a := new(Android)
  // call person struct from android
  a.Person.Talk()
  // call a person method from android struct
  a.Talk()
}
