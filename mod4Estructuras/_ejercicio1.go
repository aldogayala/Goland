package main

import (
	"fmt"
)

func main() {

	var palabra string

	fmt.Println("Ingrese una palabra")
	fmt.Scanln(&palabra)

	fmt.Printf("El largo de la palabra es: %d\n", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}

}
