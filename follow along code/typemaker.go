package main

import "fmt"

type height float64

func main() {
  var tall height
  tall = 6.3
  fmt.Printf("%T %v \n", tall, tall)
}
