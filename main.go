package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// generateRandomNumber generates an integer given the bounds of a low and high end.
// It does this immediately, but gives a bit of a show by giving a message and waiting
func generateRandomNumber(min, max int) int {
	fmt.Println("Generating Number...")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// difficultySelection asks the user to choose a difficulty through printed messages
// based on the difficulty selected.
func difficultySelection() {
	scanner := bufio.NewScanner(os.Stdin)
	var difficulty int

	for {
		fmt.Println("Select Difficulty using the corresponding number (i.e. 1 for Easy):")
		fmt.Println("1. Easy (10 Guesses, Range: 1-50)")
		fmt.Println("2. Medium (5 Guesses, Range: 1-75)")
		fmt.Println("3. Hard (3 Guesses, Range: 1-50)")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > 3 {
			fmt.Println("Invalid input. Please enter a number between 1 and 3.")
			continue
		}

		difficulty = num
		break
	}

	switch difficulty {
	case 1:
		play(generateRandomNumber(1, 50), 10)
	case 2:
		play(generateRandomNumber(1, 75), 5)
	case 3:
		play(generateRandomNumber(1, 50), 3)
	}
}

// play begins the game based on the previous difficulty and generated number choices
// Informs player on their current progress, their guess, and if they win or lose
func play(target int, totalGuesses int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Generated a number!")
	fmt.Println("You have", totalGuesses, "total guesses to guess the number.")
	fmt.Println("Good luck :^) ")
	time.Sleep(1 * time.Second)

	//Guess input begin
	for i := 0; i <= totalGuesses; i++ {
		fmt.Print("Guess: ")
		scanner.Scan()
		input := scanner.Text()

		//Guess is not a number
		//Inform them it's wrong, but don't penalize them for it
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			i-- // Don't count invalid input as an attempt
			continue
		}

		//Wildly incorrect Guess
		//Don't penalize player for it, but let em know it's wildly off
		if guess < 1 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number more reasonable.")
			fmt.Println("I'll let you off the hook for that one. You have", totalGuesses-i, "guesses remaining.")
			i-- // Don't count invalid input as an attempt
			continue
		}

		fmt.Println("You guessed", guess)
		//Warmer/Colder: Guess is off in either direction
		if guess > target {
			fmt.Println("Too high!")
		} else if guess < target {
			fmt.Println("Too low!")
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

// gameOver handles the behavior of the end state when a player wins or loses.
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

// toSnakeCase Function to capitalize the first letter of a string
// I don't know if Go had a premade function for it, so I just made it quick.
func toSnakeCase(input string) string {
	// Convert input to lowercase
	input = strings.ToLower(input)

	// Capitalize the first letter
	return strings.ToUpper(string(input[0])) + input[1:]
}
