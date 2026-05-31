package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	var result float64

	fmt.Println("Simple Calculator in Go")
	fmt.Println("========================")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Invalid operator")
		return
	}

	fmt.Println("Result:", result)
}