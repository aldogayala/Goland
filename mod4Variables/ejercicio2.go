package main

import "fmt"

func main2() {
	var temperartura float32
	var ciudad string
	var presion float32

	temperartura = 38.5
	ciudad = "Alta Gracia"
	presion = 88.3

	fmt.Printf("La tememperatura en la ciudad de %s es de %f\n", ciudad, temperartura)
	fmt.Printf("La presion en la ciudad de %s es de %f\n", ciudad, presion)
}
