### Notas
1. Existen dos maneras
 - `var x string = "Hello, World"`
 - `x := "Hello, World"`

2. `x := 5; x += 1` es igual a `6`.

3. El `scope` es el espacio en donde se define el acceso o disponibilidad de una variable o funcion. Para determinar  el `scope` de una variable, debes de ubicar en donde se encuentra colocada dicha variable, si esta se encuentra dentro de llaves `{}` significa que solo estará disponible de ahí hacía abajo y no fuera de este contexto.

4. `var` define una variable, pero esta puede ser nuevamente asignada con un valor del mismo tipo de dato, `const` no permite asignarle un nuevo valor.

5. 
```go
package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit value to convert into Celsius: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9
	fmt.Println("Celsius = ", output)
}
```

6. 
```go
package main

import "fmt"

func main() {
	fmt.Print("Enter Feet's value to convert into Meters: ")
	var input float64
	fmt.Scanf("%f", &input)
	
	output := (input * 0.3048)
	
	fmt.Println("Meters: ", output)
}
```