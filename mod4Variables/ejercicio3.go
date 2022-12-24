package main

import (
	"fmt"
	"strconv"
)

func main3() {
	var nombre string
	var apellido string
	var edad int
	edad = 28
	apellido = "Perez"
	nombre = "Juan"
	var licencia_de_conducir bool
	licencia_de_conducir = true
	var estaturaDeLaPersona float32
	estaturaDeLaPersona = 1.70

	fmt.Println("El apellido de la persona es " + apellido)
	fmt.Println("El nombre de la persona es " + nombre)
	fmt.Println("La edad de la persona es " + strconv.Itoa(edad))
	fmt.Printf("Posee licencia de conducir: %v\n", licencia_de_conducir)
	fmt.Printf("La estatura de la persona es de %f\n", estaturaDeLaPersona)

}
