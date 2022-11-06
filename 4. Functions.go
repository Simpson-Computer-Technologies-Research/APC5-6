package main

// Import fmt package
import "fmt"

// Main function
func main() {

	// Call the function that returns a value
	var returnedValue int = myFunction()

	// Call the function that takes a parameter
	myFunctionWithParameters("parameter")
}

// Return something function
func myFunction() int {
	return 1
}

// Pass a parameter to the function
func myFunctionWithParameters(param string) {
	fmt.Println(param)
}
