package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			if i % 5 == 0 {
				fmt.Println(i, "FuzzBuzz")
			} else {
				fmt.Println(i, "Fuzz")
			}
		} else if i % 5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}