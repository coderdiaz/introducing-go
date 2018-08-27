package main

import "fmt"

func main() {
	fmt.Print("Enter Feet's value to convert into Meters: ")
	var input float64
	fmt.Scanf("%f", &input)
	
	output := (input * 0.3048)
	
	fmt.Println("Meters: ", output)
}