package main

import "fmt"

func main() {
	var numeroMes int
	fmt.Println("Ingrese el numero de mes")
	fmt.Scan(&numeroMes)

	switch numeroMes {
	case 1:
		fmt.Println("El numero de mes es Enero")
	case 2:
		fmt.Println("El numero de mes es Febrero")
	case 3:
		fmt.Println("El numero de mes es Marzo")
	case 4:
		fmt.Println("El numero de mes es Abril")
	case 5:
		fmt.Println("El numero de mes es Mayo")
	case 6:
		fmt.Println("El numero de mes es Junio")
	case 7:
		fmt.Println("El numero de mes es Julio")
	case 8:
		fmt.Println("El numero de mes es Agosto")
	case 9:
		fmt.Println("El numero de mes es Septiembre")
	case 10:
		fmt.Println("El numero de mes es Octubre")
	case 11:
		fmt.Println("El numero de mes es Noviembre")
	case 12:
		fmt.Println("El numero de mes es Deciembre")
	default:
		fmt.Println("El numero ingresado es incorrecto")
	}

}
