package main

import "fmt"

func main() {

	dabble := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(dabble(96))

}
