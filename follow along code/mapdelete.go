package main

import "fmt"

func main() {

  myGreeting := map[string]string{
    "Tim":  "Good Morning!",
    "Jenny":  "Como esta",
  }

myGreeting["Harleen"] = "Howdy partner"
  fmt.Println(myGreeting)
  delete(myGreeting, "Harleen")
  fmt.Println(myGreeting)
}
