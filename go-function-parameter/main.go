package main

import "fmt"

func getName(nickName string) { // Function to print the given string
	fmt.Println("Hallo", nickName) // Print the string
}

func getPeople(myName string, myAge int) { // Function with multiple parameters
	fmt.Println("My name is", myName, "and I am", myAge, "years old.") // Print the values of the parameters
}

func main() {
	getName("Rifqy") // Call the function with an argument

	getPeople("Muhammad Rifqy", 20) // Call the function with multiple arguments
}