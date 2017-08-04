package main

import "fmt"

func dabble(n int) (int, bool) {
  return n / 2, n%2 == 0
}

func main() {
  h, even := dabble(12)
  fmt.Println(h, even)
}
