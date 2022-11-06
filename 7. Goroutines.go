package main

// Import packages
import (
	"fmt"
	"time"
)

// Main function
func main() {
	// Call the functions in a goroutine
	go myGoroutineFunction()
	go myGoroutineFunction()

	// Create a new function and call it like so
	go func() {
		fmt.Println(time.Now())
	}()

	// This will return an error.
	// Goroutines CAN NOT returns values like normal functions
	// Define: "return ..." does not work
	var returnedValue = go myGoroutineFunction()
}

// Goroutines are like threads. You can call multiple functions
// at the same time instead of one after the other
func myGoroutineFunction() int {

	// Print the current time
	fmt.Println(time.Now())

	// Return a temp value
	return 1
}
