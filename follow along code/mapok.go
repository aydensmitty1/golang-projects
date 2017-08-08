package main

import "fmt"

func main() {

  myGreeting := map[int]string{
  0: "Good Morning",
  1: "Buenos Dias",
  2: "Bonjour",
  3: "Buongiorno",
}

fmt.Println(myGreeting)
//delete(myGreeting, 2)
if val, exists := myGreeting[2]; exists {
  fmt.Println("val: ", val)
  fmt.Println("exists: ", exists)
  } else {
    fmt.Println("That value doesn't exist.")
    fmt.Println("val: ", val)
    fmt.Println("exists: ", exists)

  }
  fmt.Println(myGreeting)
}
