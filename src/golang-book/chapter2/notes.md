### Notes
1. Las computadoras almacenan los números enteros en un sistema binario de Base 2.
2. 1 2 4 6 8 16 32 64 = 11111111 = 133 en Base 10
3. 
```go
package main

import "fmt"

func main () {
	fmt.Println("32132 * 42452 =", 32132 * 42452)
}
```
4. Los tipos `strings` son una secuencia de carácteres con un tamaño definido y que representa un texto. Para nosotros conocer el tamaño de la cadena hacemos uso del helper `len()`.
5. La respuesta es: `false`.
```go
package main

import "fmt"

func main () {
	fmt.Println((true && false ) || (false && true) || !(false && false))
}
```