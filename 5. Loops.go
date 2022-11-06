package main

// Import fmt package
import "fmt"

// Main function
func main() {

	// Ranged loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While / Conditional Loop
	for {
		fmt.Println(2)
	}

	// While 1 == 1
	for 1 == 1 {
		fmt.Println(3)
	}

	// Declare a map for testing
	var myMap = map[string]int{"key": 1}

	// Iterating over a map
	for key, value := range myMap {
		fmt.Printf("Key: %s\nValue: %s\n\n", key, value)
	}
}
