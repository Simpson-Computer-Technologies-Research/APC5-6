package main

// Import fmt package
// YOU MUST USE THE FMT PACKAGE
//
// Luckily it's provided by Golang so you don't need to
// install anything
import "fmt"

// Main function
func main() {

	// Print text on a new line
	fmt.Println("Hello World!")

	// Merge Strings
	fmt.Println("My name is: " + "Ronald")

	// Formatting
	// %s for strings
	// %d for integers
	// %f for floats
	// %v for any
	fmt.Printf("Name: %s\nAge: %d\n\n", "tristan", 17)

	// Declare an empty variable.
	// I'll explain more about variables in the Variables.go lesson
	var userInput string

	// Print to terminal what the user must input
	fmt.Print("Your Full Name is: ")
	// Use Scanln to take in the input
	fmt.Scanln(&userInput)
}
