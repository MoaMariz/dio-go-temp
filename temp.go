package main

import "fmt"

func main() {

	var tempK float64
	fmt.Println("Digite a temperatura em graus Kelvin")
	fmt.Scanln(&tempK)

	var tempC = tempK - 273
	fmt.Println("A temperatura convertida em Celsius =", tempC, "ºC")

}
