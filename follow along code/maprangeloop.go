package main

import "fmt"

func main() {

  myGreeting := map[int]string{
  0: "Good Morning",
  1: "Buenos Dias",
  2: "Bonjour",
  3: "Buongiorno",
  }

    for key, val := range myGreeting {
      fmt.Println(key, " - ", val)
    }
  }
