package main

import (
        "math/rand"
        "fmt"
        "time"
)

func main() {
  again := rand.NewSource(time.Now().UnixNano())
  next := rand.New(again)
  fmt.Println("\n Here is your Magic Number")
  fmt.Println(
    next.Intn(6))

}
