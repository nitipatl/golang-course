package main

import "fmt"

func main() {
	var fahrenheit float64
	var celsius float64
	fmt.Print("Enter feet: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Println(celsius)
}
