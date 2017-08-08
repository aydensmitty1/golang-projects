package main

import (
        "fmt"
        "sort"
)

func main() {
s := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
fmt.Println(s)
sort.Ints(s)
fmt.Println(s)
}
