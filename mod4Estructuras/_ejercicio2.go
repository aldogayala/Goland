package main

import "fmt"

func main() {
	var edad int
	fmt.Println("Ingrese cuantos años posee el cliente")
	fmt.Scan(&edad)
	if edad > 22 {

		var empleado bool
		fmt.Println("¿El clientes esta trabajandoa actualmente? True/False")
		fmt.Scan(&empleado)
		if empleado {
			var añoAntiguedad int
			fmt.Println("Ingrese años de antiguedad del cliente")
			fmt.Scan(&añoAntiguedad)
			if añoAntiguedad > 1 {
				var sueldo float64
				fmt.Println("Ingrese sueldo del cliente")
				fmt.Scan(&sueldo)
				fmt.Println("Se otorga el prestamo al cliente")
				if sueldo > 100000 {
					fmt.Println("No se cobra comision al prestamo")
				}
			} else {
				fmt.Println("El cliente debe poseer mas de un año de antiguedad")
			}

		} else {
			fmt.Println("No se puede otorgar el prestamo")
			fmt.Println("El cliente debe poseer un trabajo")
		}

	} else {
		fmt.Println("No se puede otorgar el prestamos")
		fmt.Println("El cliente debe tener mas de 22 años")
	}

}
