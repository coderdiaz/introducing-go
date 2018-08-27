package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	
	// fmt.Println(elements["Un"]) // Getting a not existed element, this return empty string
	
	// name, ok := elements["Un"] // Go helper for know if element exists, if exists return the value in name
	// fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}