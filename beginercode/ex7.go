package main

import "fmt"

func main() {
  three := 0
  for i := 0; i < 1000; i++ {
    if i%3 == 0 {
    three += i
  } else if i%5 == 0 {
    three += i
}
}
    fmt.Println(three)
}
