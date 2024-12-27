package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

//@Author: <John Botonakis>
//@Date: 2021-03-29
//@Description: This is a number guessing game where the player is asked to guess a number.
//The amount of guesses will vary from difficulty level.

var playing = true

// Main function holding the beginning instructions and prompts for beginning the game
func main() {
	//TO DO: Input Validation - make sure it's alpha and not alphanumeric
	fmt.Printf("Please Enter your Name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Printf("Hello %v! Welcome to the number guessing game!\n", name)
	fmt.Println("The goal of the game is to guess the number I'm thinking of.")
	fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")
	for playing {
		fmt.Printf("Select Difficulty:\n1. Easy (5 Guesses, 1-50) \n2. Medium (5 Guesses, 1-75) \n3. Hard (3 Guesses, 1- 33)\n")

		//Select Difficulty
		var difficulty int
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			return
		}
		//Difficulty Selection: Each difficulty will have their own max and amount of guesses
		switch difficulty {
		case 1:
			fmt.Printf("Difficulty Selected: %v Easy (5 Guesses)\n", difficulty)
			easyNumber := randomNumber(1, 50)
			play(easyNumber, 5)
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
func play(target int, totalguesses int) {
	fmt.Println("Generated a number!")
	fmt.Println("You have", totalguesses, "total guesses to guess the number.")
	fmt.Println("Good luck :^) ")
	time.Sleep(1 * time.Second)

	//Guess input begin
	for i := 0; i < totalguesses; i++ {
		fmt.Print("Guess: ")
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			return
		}
		fmt.Println("You guessed", guess)

		//Warmer/Colder: Guess is off by 15 in either direction
		if guess > target && guess <= target+15 {
			fmt.Println("Warmer!")
		} else if guess < target && guess >= target-15 {
			fmt.Println("Warmer!")
		} else if guess > target+15 || guess < target-15 {
			fmt.Println("Colder!")
		}

		//Last guess notification
		if i == totalguesses-2 {
			fmt.Println("Last guess! Here's hoping you get it right!")
			time.Sleep(1 * time.Second)
		}
		//Win Condition
		if guess == target {
			winner(target, i)
			return
		}

		if i == totalguesses-1 {
			fmt.Printf("You ran out of guesses. The correct number was %v !", target)
			game_over()
		}
	}
}

// Winning Condition is met
func winner(correctNumber int, guesses_left int) {
	fmt.Printf("You win! You got the right number: %v in %v guesses!", correctNumber, guesses_left)
}

func game_over() {
	fmt.Println("Game Over!")
	fmt.Println("Would you like to play again? 1. Yes, 2. No")

	//Play Again input and validation
	var playAgain int
	_, err := fmt.Scanln(&playAgain)
	if err != nil {
		return
	}

	if playAgain == 1 {
		return //Continue the Loop uninterrupted
	}
	if playAgain == 2 {
		playing = false //Break the Loop
	}
}
