package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("La edad de Benjamin es: %d\n", employees["Benjamin"])

	contador := 0

	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Printf("Cantidad de empleados con edad > 21: %d\n", contador)
	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}
