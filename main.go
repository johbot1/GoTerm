package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

//@Author: <John Botonakis>
//@Date: 2021-03-29
//@Description: This is a number guessing game where the player is asked to guess a number.
//The amount of guesses will vary from difficulty level.

//TODO:fmt.Scan(&difficulty) does not handle invalid input
//TODO:If a user enters non-numeric input, fmt.Scanln(&guess) fails without any noise
//TODO: Guess check (guess > 100 || guess < 0) resets the guess count incorrectly.

var playing = true

// Main function holding the beginning instructions and prompts for beginning the game
func main() {
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	// Loop until a valid name is entered
	for {
		fmt.Printf("Please Enter your Name (No spaces, only letters please): ")

		// Read and store the input line
		scanner.Scan()
		name = scanner.Text()

		// Trim spaces to avoid accidental leading/trailing spaces
		name = strings.TrimSpace(name)

		// Validate input using the updated validateName function
		if !validateName(name) {
			// Provide feedback and retry if validation fails
			continue
		}

		// Convert and format valid input
		name = toSnakeCase(name)
		break
	}

	//How to Play instructions
	fmt.Printf("Hello %v! Welcome to the number guessing game!\n", name)
	//Main Loop; When finishing a game, return to THIS point

	for playing {
		fmt.Println("The goal of the game is to guess the number I'm thinking of.")
		fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")
		time.Sleep(2 * time.Second)
		difficultySelection()
		if !playing {
			break
		}
	}
	fmt.Println("Thank you for playing!")
}

// randomNumber generates a random integer between the given min and max values (inclusive).
// Generated with JetBrains AI code commenting
func randomNumber(min, max int) int {
	fmt.Println("Generating Number...")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// difficultySelection prompts the user to select a difficulty level and initiates the number-guessing game accordingly.
// Generated with JetBrains AI code commenting
func difficultySelection() {
	var difficulty int
	for {
		fmt.Println("Select Difficulty using the corresponding number (i.e. 1 for Easy):")
		fmt.Println("1. Easy (10 Guesses, Range: 1-50)")
		fmt.Println("2. Medium (5 Guesses, Range: 1-75)")
		fmt.Println("3. Hard (3 Guesses, Range: 1-50)")

		//Validate Input
		_, err := fmt.Scan(&difficulty)
		if err != nil || difficulty > 3 || difficulty < 1 {
			// Handle non-integer input
			fmt.Println("Invalid input. Please enter a number between 1 and 3.")
			// Clear any input
			var discard string
			_, _ = fmt.Scanln(&discard)
			continue
		}

		//Difficulty Selection: Each difficulty will have their own max and amount of guesses
		switch difficulty {
		case 1:
			fmt.Printf("Difficulty Selected: %v Easy (5 Guesses)\n", difficulty)
			easyNumber := randomNumber(1, 50)
			play(easyNumber, 10)
		case 2:
			fmt.Printf("Difficulty Selected: %v Medium (5 Guesses)\n", difficulty)
			mediumNumber := randomNumber(1, 75)
			play(mediumNumber, 5)
		case 3:
			fmt.Printf("Difficulty Selected: %v Hard (3 Guesses)\n", difficulty)
			hardNumber := randomNumber(1, 50)
			play(hardNumber, 3)
		default:
			fmt.Println("Invalid selection. Please choose 1, 2, or 3.")
		}
		break
	}
}

// play initiates a number-guessing game where the user attempts to guess the target within a limited number of guesses.
// Generated with JetBrains AI code commenting
func play(target int, totalGuesses int) {
	fmt.Println("Generated a number!")
	fmt.Println("You have", totalGuesses, "total guesses to guess the number.")
	fmt.Println("Good luck :^) ")
	time.Sleep(1 * time.Second)

	//Guess input begin
	//TODO: Validation for super high guesses
	for i := 0; i <= totalGuesses; i++ {
		fmt.Print("Guess: ")
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			return
		}
		fmt.Println("You guessed", guess)

		//Warmer/Colder: Guess is off in either direction
		if guess > target {
			fmt.Println("Too high!")
		} else if guess < target {
			fmt.Println("Too low!")
		}

		//Wildly incorrect Guess
		//Don't penalize player for it, but let em know it's wildly off
		if guess > 100 || guess < 0 {
			if i > 0 {
				i = i - 1
			}
			fmt.Println("Invalid input. Please enter a number more reasonable.")
			fmt.Println("I'll let you off the hook for that one. You have", totalGuesses-i, "guesses remaining.")
			continue
		}

		//Win Condition
		if guess == target {
			gameOver(target, i, true)
			break
		}

		//Last guess notification + Game Over Condition
		if i == totalGuesses-2 {
			fmt.Println("Last guess! Here's hoping you get it right!")
			time.Sleep(1 * time.Second)
		} else if i == totalGuesses-1 {
			fmt.Printf("You ran out of guesses!")
			gameOver(target, i, false)
			break
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

	//Play Again input and validation
	fmt.Printf("\nWould you like to play again?\n1. Yes, 2. No ")
	var playAgain int
	for {
		_, err := fmt.Scanln(&playAgain)
		if err != nil {
			return
		} else if playAgain < 1 || playAgain > 2 {
			fmt.Println("Invalid input. Please enter 1 or 2.")
		} else {
			break
		}
	}
	switch playAgain {
	case 1:
		fmt.Println("The grind never stops")
	case 2:
		playing = false // Exit the main loop
		fmt.Println("Thank you for playing!! Goodbye!")
		break
	}
}

// Function to validate if a string contains alphabetical characters only
func validateName(input string) bool {
	// Check for empty strings
	if strings.TrimSpace(input) == "" {
		fmt.Println("Invalid input. Name cannot be empty.")
		return false
	}

	// Check for spaces
	if strings.Contains(input, " ") {
		fmt.Println("Invalid input. Name cannot contain spaces.")
		return false
	}
	//Check against the alphabet in lower and upper case
	pattern := "^[A-Za-z]+$"
	//CODE NOTE: The "^" indicates the start of the string, "$" denotes the end of the string

	match, err := regexp.MatchString(pattern, input)

	//Check to see if there is any error in the matching. If there is, throw it.
	if err != nil {
		fmt.Println("Error occurred during validateName:", err)
		return false
	}
	return match
}

// Function to capitalize the first letter of a string
func toSnakeCase(input string) string {
	// Convert input to lowercase
	input = strings.ToLower(input)

	// Capitalize the first letter
	return strings.ToUpper(string(input[0])) + input[1:]
}
