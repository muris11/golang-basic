package main

import "fmt"

func greet() { // Function without parameters
	fmt.Println("Hello, World!") // Print a greeting message
}

func add(a int, b int) int { // Function with parameters and return value
	return a + b // Return the sum of a and b
}

func main() {
	greet() // Call the function without arguments

	result := add(5, 3) // Call the function with arguments and store the result
	fmt.Println("5 + 3 =", result) // Print the result
}