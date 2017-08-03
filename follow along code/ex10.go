package main

import "fmt"

func max(sf ...int) int {
  var largest int
  for _, v := range sf {
    if v > largest {
      largest = v
    }
  }
return largest
}
func main() {
  n := max(11, 15, 17 ,19, 22)
  fmt.Println(n)
}
