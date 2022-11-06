package main

// Import fmt package
import "fmt"

// Create a new data structure named "Person"
type Person struct {
	name string
	age  int
}

// Create a new variable
type NewVariable []byte

// Main function
func main() {

	// Declare a new variable with your new type
	var byteArray NewVariable = []byte{1, 1, 1}

	// Create a new Person Object
	var user Person = Person{
		name: "Tristan",
		age:  17,
	}

	// Print the data
	fmt.Println(user.name)
	fmt.Println(user.age)
}
