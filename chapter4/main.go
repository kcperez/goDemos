package main

import "fmt"

//Using the example program as a starting point,
//write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)

func main() {
	fmt.Print("Convert Fahrenheit into Celcius: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5 / 9
	fmt.Println(output)
}
