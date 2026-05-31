package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string
	var result float64
	var choice string

	fmt.Println("Simple Calculator in Go")
	fmt.Println("========================")

	// Keep the calculator running until the user decides to stop
	for {

		// Ask for the first number
		fmt.Print("\nEnter first number: ")
		fmt.Scanln(&num1)

		// Ask what operation to perform (or quit)
		fmt.Print("Enter operator (+, -, *, /) or q to quit: ")
		fmt.Scanln(&operator)

		// Let user exit anytime they want
		if operator == "q" || operator == "Q" {
			fmt.Println("Exiting calculator... see you next time 👋")
			break
		}

		// Get the second number for the calculation
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		// Do the math based on the operator chosen
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			// Prevent crashing when dividing by zero
			if num2 == 0 {
				fmt.Println("Error: You can't divide by zero")
				continue
			}
			result = num1 / num2
		default:
			// Catch anything that isn't a valid operator
			fmt.Println("Error: That operator doesn't exist here")
			continue
		}

		// Show the final answer
		fmt.Println("Result:", result)

		// Ask if the user wants to run another calculation
		fmt.Print("Do another calculation? (y/n): ")
		fmt.Scanln(&choice)

		if choice != "y" && choice != "Y" {
			fmt.Println("Goodbye 👋")
			break
		}
	}
}