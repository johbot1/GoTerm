package main

import "fmt"

//@Author: <John Botonakis>
//@Date: 2021-03-29
//@Description: This is a number guessing game where the player is asked to guess a number.
//The amount of guesses will vary from difficulty level.

func main() {
	fmt.Printf("Please Enter your Name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v! Welcome to the number guessing game!", name)
	fmt.Println("The goal of the game is to guess the number I'm thinking of.")
	fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")

	//Select Difficulty
	fmt.Printf("Select Difficulty:\n1. Easy (5 Guesses) \n2. Medium (5 Guesses) \n3. Hard (3 Guesses)\n")
	var difficulty uint
	fmt.Scan(difficulty)

}
