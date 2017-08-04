package main

import "fmt"

func main() {
m := make(map[string]int)
m["k1"] = 7
m["k2"] = 13
//n := map[string]int{"Foo": 1, "Bar": 2}
  //delete(m, "k2")
  fmt.Println(m)

  v, ok := m["k2"]
  fmt.Println("ok?", ok, v)
}
