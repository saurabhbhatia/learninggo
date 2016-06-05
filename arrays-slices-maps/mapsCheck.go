package main

import "fmt"

func main() {
  // map definition
  x := make(map[string]int)
  x["1"] = 10
  fmt.Println(x["1"])

  // map definition with multiple elements
  mapElements := make(map[string]string)
  mapElements["H"] = "Hydrogen"
  mapElements["He"] = "Helium"
  mapElements["O"] = "Oxygen"
  mapElements["Li"] = "Lithium"

  // display entire map
  fmt.Println(mapElements)
  // display an element
  fmt.Println(mapElements["Li"])

  //multi key map elements
  multiElements := map[string]map[string]string{
    "H": map[string]string {
      "name": "Hydrogen",
      "state": "gas",
    },
    "He": map[string]string {
      "name": "Helium",
      "state": "gas",
    },
    "Ca": map[string]string {
      "name": "Calcium",
      "state": "solid",
    },
  }
  if el, ok := multiElements["Ca"]; ok {
    fmt.Println(el["name"], el["state"])
  }
}
