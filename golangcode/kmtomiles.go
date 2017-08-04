package main

import "fmt"

const KilometersToMiles float64 = 0.621371

func main() {
var Kilometers float64
fmt.Print("Enter Kilometers ran: ")
fmt.Scan(&Kilometers)
Miles := Kilometers * KilometersToMiles
fmt.Println(Kilometers, " You ran ", Miles, " miles. Congratulations")
}
