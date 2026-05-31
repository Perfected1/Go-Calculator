package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator in Go")
	fmt.Println("========================")

	// Take first number
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	// Take operator
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	// Take second number
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Processing...")
}