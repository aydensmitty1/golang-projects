package main

import "fmt"

type person struct {
  First string
  Last string
  Age int
}
func main() {
  p1 := person{"James", "Bond", 20}
  p2 := person{"Kodac", "Black", 25}
  fmt.Println(p1.First, p1.Last, p1.Age)
  fmt.Println(p2.First, p2.Last, p2.Age)
}
