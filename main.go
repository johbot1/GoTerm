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
	var name string
	//Loop until a valid name is entered
	for {
		fmt.Printf("Please Enter your Name (alphabetical only): ")
		_, err := fmt.Scanln(&name) // Read the input
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Validate input using regex
		isValid := validateAlphabetical(name)
		if !isValid {
			fmt.Println("Invalid input. Please enter letters only, without spaces or special characters.")
			continue
		}

		// Valid input
		break
	}

	fmt.Printf("Hello %v! Welcome to the number guessing game!\n", name)
	fmt.Println("The goal of the game is to guess the number I'm thinking of.")
	fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")

	//Difficulty Selection and input validation
	for playing {
		fmt.Printf("Select Difficulty:\n1. Easy (5 Guesses, 1-50) \n2. Medium (5 Guesses, 1-75) \n3. Hard (3 Guesses, 1- 33)\n")

		//Select Difficulty
		var difficulty int
		_, err := fmt.Scanln(&difficulty)
		if err != nil {
			return
		} else if difficulty < 1 || difficulty > 3 {
			fmt.Println("Invalid input. Please enter 1, 2, or 3.")
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
func play(target int, totalGuesses int) {
	fmt.Println("Generated a number!")
	fmt.Println("You have", totalGuesses, "total guesses to guess the number.")
	fmt.Println("Good luck :^) ")
	time.Sleep(1 * time.Second)

	//Guess input begin
	for i := 0; i < totalGuesses; i++ {
		fmt.Print("Guess: ")
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			return
		}
		fmt.Println("You guessed", guess)

		//Warmer/Colder: Guess is off by 10 in either direction
		if guess > target && guess <= target+10 {
			fmt.Println("Warmer!")
		} else if guess < target && guess >= target-10 {
			fmt.Println("Warmer!")
		} else if guess > target+10 || guess <= target-10 {
			fmt.Println("Colder!")
		}

		//Win Condition
		if guess == target {
			gameOver(target, i, true)
			return
		}

		//Last guess notification + Game Over Condition
		if i == totalGuesses-2 {
			fmt.Println("Last guess! Here's hoping you get it right!")
			time.Sleep(1 * time.Second)
		} else if i == totalGuesses-1 {
			fmt.Printf("You ran out of guesses! The correct number was %v.", target)
			gameOver(target, i, false)
		}
	}
}

// gameOver handles the end of the game by prompting the user to decide whether to play again or terminate the game.
// Generated with JetBrains AI code commenting
func gameOver(correctNumber int, guessesLeft int, win bool) {
	if win {
		fmt.Printf("You win! You got the right number: %v in %v guesses!", correctNumber, guessesLeft+1)
	} else {
		fmt.Printf("\nGame Over! Womp womp! The correct number was %v.", correctNumber)
	}
	fmt.Printf("\nWould you like to play again?\n1. Yes, 2. No ")

	//Play Again input and validation
	var playAgain int
	_, err := fmt.Scanln(&playAgain)
	if err != nil {
		return
	} else if playAgain < 1 || playAgain > 2 {
		fmt.Println("Invalid input. Please enter 1 or 2.")
	}

	if playAgain == 1 {
		return //Continue the Loop uninterrupted
	}
	if playAgain == 2 {
		playing = false //Break the Loop
	}
}

// Function to validate if a string contains alphabetical characters only
func validateAlphabetical(input string) bool {
	//Check against the alphabet in lower and upper case
	pattern := "^[A-Za-z]+$"
	//CODE NOTE: The "^" indicates the start of the string, "$" denotes the end of the string

	match, err := regexp.MatchString(pattern, input)

	//Check to see if there is any error in the matching. If there is, throw it.
	if err != nil {
		fmt.Println("Error occurred during validateAlphabetical:", err)
		return false
	}
	return match
}
