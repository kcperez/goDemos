package main

import "fmt"

//Write another program that converts from feet into meters. (1 ft = 0.3048 m)

func main() {
	fmt.Println("Convert feet into meters: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output)
}
