package main

import "fmt"

func main() {
  foo(1, 2)
  foo(1, 2, 3)
  aslice := []int{1, 2, 3, 4}
  foo(aslice...)
  foo()
}

func foo(numbers ...int) {
  fmt.Println(numbers)
}
