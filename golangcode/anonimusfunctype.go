package main

import "fmt"

func main() {
  greeting := func() {
    fmt.Println("hello world")
  }

  greeting()
  fmt.Printf("%T\n", greeting)
}
