### Notes
1. Un espacio en blanco entendido por el lenguaje como un salto de linea.
2. Un comentario, es la forma en la que puedes documentar tu código para que alguien más pueda interpretar tu código mucho más rápido. Hay dos formas de utilizarlo:
  - Con `//`, este comenta por línea.
  - Y, con '/* */', este comenta en multiples líneas.
3. Este iniciaría con `package fmt`.
4. Debemos importar el paquete `os` de Go para poder hacer uso de la función `Exit`, esta recíbe un parámetro `int`, donde `0` es `sucess` y algún otro número es `error`
```go
package main

import "fmt"
import "os"

// This is a comment

func main() {
	fmt.Println("Hello World")
	os.Exit(0)
}
```

5.
```go
package main

import "fmt"

// This is a comment

func main() {
	fmt.Println("Hello, my name is Javier")
}
```