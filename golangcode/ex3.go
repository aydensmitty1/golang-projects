package main

import "fmt"

func main() {
	var name string
	var feelings string
	fmt.Println("Please Enter your name")
	fmt.Scan(&name)
	fmt.Println("Hello", name, "how are you?")
	fmt.Scan(&feelings)
	switch feelings {
	case "great", "good", "Great", "Good", "Awsome", "awsome":
		fmt.Println("Great, lets start")
	case "bad", "awfull", "sucky":
		fmt.Println("well thats too bad")
		fallthrough
	default:
		fmt.Println("your day is about to get even lot better")
	}
}
