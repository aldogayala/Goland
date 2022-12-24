package main

import (
	"fmt"
	"strconv"
)

func main() {
	apellido := "Gomez"
	edad := 35
	flag := false
	sueldo := 45857.90
	nombre := "Julián"

	fmt.Println("El nombre del empleado es " + apellido + " " + nombre)
	fmt.Println("La edad del empleado es de: " + strconv.Itoa(edad))
	fmt.Printf("El valor de boolean es: %v\n", flag)
	fmt.Printf("El sueldo del empleado es: %f\n", sueldo)

}
