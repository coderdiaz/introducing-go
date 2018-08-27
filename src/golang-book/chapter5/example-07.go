package main

import "fmt"

func main() {
	var x map[string]int // The maps have to be intialized before they can be used
	x["key"] = 20
	fmt.Println(x)
}