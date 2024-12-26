// @Author: <John Botonakis>
//
//	@Date: 2021-03-29
//
// @Description: This is a cheat sheet for Go, helping me undertstand this new languag. This has been put together by
// myself and Google Gemini for my own reference.
package main

import (
	"bufio"   // Package for buffered I/O operations, allowing efficient reading of input from the console.
	"fmt"     // Package for formatted I/O, providing functions like Println for printing to the console.
	"os"      // Package for operating system-level functions, including access to standard input (os.Stdin).
	"strconv" // Package for string conversions, used to convert strings to numbers and vice versa.
	"strings" // Package for string manipulation, providing functions like TrimSpace for removing whitespace.
)

// This program demonstrates basic Go concepts for terminal projects.
func main() {
	// --- Data Types ---

	//// Strings: Sequence of characters.
	//name := "John Doe"
	//fmt.Println("Name:", name)
	//
	//// Integers: All whole numbers, negatives included
	//age := 30
	//fmt.Println("Age:", age)
	//
	//// Unsigned integers: Non-negative whole numbers.
	//id := uint(12345) // Explicitly cast to unsigned int
	//fmt.Println("ID:", id)
	//
	//// Floating-point numbers (double-precision).
	//price := 99.99
	//fmt.Println("Price:", price)
	//
	//// Booleans: True or false values.
	//isStudent := true
	//fmt.Println("Is Student:", isStudent)

	// --- User Input ---

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	userInput, _ := reader.ReadString('\n')  // Read input until newline
	userInput = strings.TrimSpace(userInput) // Remove leading/trailing whitespace
	fmt.Println("You entered:", userInput)

	// --- String Manipulation ---

	// Concatenation
	name := userInput
	greeting := "Hello, " + name + "!"
	fmt.Println(greeting)

	// String formatting (using fmt.Sprintf)
	formattedGreeting := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(formattedGreeting)

	// --- Input Conversion ---

	fmt.Print("Enter your age: ")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)

	// Convert string to integer
	age, err := strconv.Atoi(ageInput)
	if err != nil {
		fmt.Println("Invalid age input.")
		return
	}
	fmt.Println("Your age is:", age)

	// --- Conditional Statements ---

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// --- Loops ---

	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// --- Arrays and Slices ---

	// Array (fixed size)
	numbers := [5]int{1, 2, 3, 4, 5}
	// Print numbers sequentially
	fmt.Println("Numbers:", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])

	// Slice (dynamic size)
	fruits := []string{"apple", "banana", "orange"}
	// Print only "Banana" from the fruits slice
	fmt.Println("Fruit:", fruits[1])

	// --- Functions ---
	result := add(5, 3)
	fmt.Println("Result:", result)
}

// Function to add two integers
func add(a, b int) int {
	return a + b
}
