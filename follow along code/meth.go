package main

import "fmt"

type person struct {
  First string
  Last string
  Age int
}

func (p person) FullName() string {
  return p.First + p.Last
}
func main() {
  p1 := person{"James", "Bond", 20}
  p2 := person{"Peter", "Parker", 25}
  fmt.Println(p1.FullName())
  fmt.Println(p2.FullName())
}
