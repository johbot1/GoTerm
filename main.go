package main

import (
	"fmt"
	"math/rand"
	"time"
)

//@Author: <John Botonakis>
//@Date: 2021-03-29
//@Description: This is a number guessing game where the player is asked to guess a number.
//The amount of guesses will vary from difficulty level.

// Main function holding the beginning instructions and prompts for beginning the game
func main() {
	fmt.Printf("Please Enter your Name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %v! Welcome to the number guessing game!\n", name)
	fmt.Println("The goal of the game is to guess the number I'm thinking of.")
	fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")
	var playing bool = true
	for playing {
		fmt.Printf("Select Difficulty:\n1. Easy (5 Guesses) \n2. Medium (5 Guesses) \n3. Hard (3 Guesses)\n")

		//Select Difficulty
		var difficulty int
		fmt.Scan(&difficulty)

		switch difficulty {
		case 1:
			fmt.Printf("Difficulty Selected: %v Easy (5 Guesses)\n", difficulty)
			easyNumber := randomNumber(1, 100)
			fmt.Println(easyNumber)
		case 2:
			fmt.Printf("Difficulty Selected: %v Medium (5 Guesses)\n", difficulty)
			mediumNumber := randomNumber(1, 100)
			fmt.Println(mediumNumber)
		case 3:
			fmt.Printf("Difficulty Selected: %v Hard (3 Guesses)\n", difficulty)
			hardNumber := randomNumber(1, 100)
			fmt.Println(hardNumber)
		default:
			fmt.Println("Invalid selection. Please choose 1, 2, or 3.")
		}

	}
}

// randomNumber generates a random integer between the given min and max values (inclusive).
// Generated with JetBrains AI code commenting
func randomNumber(min, max int) int {
	fmt.Println("Generating Number...")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// play initiates a number-guessing game where the user attempts to guess the target within a limited number of guesses.
// Generated with JetBrains AI code commenting
func play(target int, guesses int) {

}
