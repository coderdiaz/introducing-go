### Notes

1. `small`.

2. 
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i ,"Divisible by 3")
		}
	}
}
```

3. 
```go
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
```