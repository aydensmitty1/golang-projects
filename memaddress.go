package main

import "fmt"

func main() {

    a := 43

    fmt.Println("a - ", a)
    fmt.Println("a's memory adress - ", &a)
    fmt.Printf("%d", &a)
}
