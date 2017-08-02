package main

import "fmt"

func main() {
  var big int
  var small int
  fmt.Println("please enter a double or tripple diget whole number")
  fmt.Scan(&big)
  fmt.Println("now enter a single diget whole number")
  fmt.Scan(&small)
  dividend := big/small
  remainder := big%small
  fmt.Println("Answer:", dividend, "with a remainder of", remainder)

}
