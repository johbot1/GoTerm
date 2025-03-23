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

const (
	EasyDifficulty   = 1
	MediumDifficulty = 2
	HardDifficulty   = 3

	EasyGuesses   = 10
	MediumGuesses = 5
	HardGuesses   = 3

	EasyMinRange   = 1
	EasyMaxRange   = 50
	MediumMinRange = 1
	MediumMaxRange = 75
	HardMinRange   = 1
	HardMaxRange   = 50

	DifficultyOptions = 3

	PlayAgainYes = "Y"
	PlayAgainNo  = "N"

	GeneratingNumberDelay = 2 * time.Second
	FeedbackDelay         = 1 * time.Second
)

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
			continue
		}

		// Convert and format valid input
		name = toSnakeCase(name)
		break
	}

	//How to Play instructions
	fmt.Printf("Hello %v! Welcome to the number guessing game!\n", name)
	fmt.Println("The goal of the game is to guess the number I'm thinking of.")
	fmt.Println("Each difficulty gives you a different amount of chances to guess my number. Good luck!")
	//Main Loop; When finishing a game, return to THIS point
	for playing {
		difficultySelection()
		if !playing {
			break
		}
	}
}

// generateRandomNumber generates an integer given the bounds of a low and high end.
// It does this immediately, but gives a bit of a show by giving a message and waiting
func generateRandomNumber(min, max int) int {
	fmt.Println("Generating Number...")
	time.Sleep(GeneratingNumberDelay)
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
		fmt.Printf("%d. Easy (%d Guesses, Range: %d-%d)\n", EasyDifficulty, EasyGuesses, EasyMinRange, EasyMaxRange)
		fmt.Printf("%d. Medium (%d Guesses, Range: %d-%d)\n", MediumDifficulty, MediumGuesses, MediumMinRange, MediumMaxRange)
		fmt.Printf("%d. Hard (%d Guesses, Range: %d-%d)\n", HardDifficulty, HardGuesses, HardMinRange, HardMaxRange)
		fmt.Print("Enter choice: ")

		scanner.Scan()
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil || num < EasyDifficulty || num > DifficultyOptions {
			fmt.Println("Invalid input. Please enter a number between", EasyDifficulty, "and", DifficultyOptions, ".")
			continue
		}

		difficulty = num
		break
	}

	switch difficulty {
	case EasyDifficulty:
		play(generateRandomNumber(EasyMinRange, EasyMaxRange), EasyGuesses, EasyDifficulty)
	case MediumDifficulty:
		play(generateRandomNumber(MediumMinRange, MediumMaxRange), MediumGuesses, MediumDifficulty)
	case HardDifficulty:
		play(generateRandomNumber(HardMinRange, HardMaxRange), HardGuesses, HardDifficulty)
	}
}

// play begins the game based on the previous difficulty and generated number choices
// Informs player on their current progress, their guess, and if they win or lose
func play(target int, totalGuesses int, difficulty int) {
	var difficultyCeiling int
	switch difficulty {
	case EasyDifficulty:
		difficultyCeiling = EasyMaxRange
	case MediumDifficulty:
		difficultyCeiling = MediumMaxRange
	case HardDifficulty:
		difficultyCeiling = HardMaxRange
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Generated a number!")
	fmt.Println("You have", totalGuesses, "total guesses to guess the number.")
	fmt.Println("Good luck :^) ")
	time.Sleep(FeedbackDelay)

	//Guess input begin
	for num_guesses := 0; num_guesses < totalGuesses; num_guesses++ {
		fmt.Print("Guess: ")
		scanner.Scan()
		input := scanner.Text()

		//Guess is not a number
		//Inform them it's wrong, but don't penalize them for it
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number for your guess.")
			continue
		}

		//Wildly incorrect Guess
		//Don't penalize player for it, but let em know it's wildly off
		if guess < EasyMinRange || guess > difficultyCeiling {
			fmt.Printf("Invalid input. Please enter a number between %d and %d.\n", EasyMinRange, difficultyCeiling)
			fmt.Printf("I'll let you off the hook for that one. You have %d guesses remaining.\n", totalGuesses-num_guesses)
			num_guesses-- // Don't count invalid input as an attempt
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
			gameOver(target, num_guesses, true)
			return // Exit the play function
		}
		//Last guess notification + Game Over Condition
		if num_guesses == totalGuesses-1 {
			fmt.Println("Last guess! Here's hoping you get it right!")
			time.Sleep(FeedbackDelay)
		}
	}
	// Ran out of guesses
	fmt.Printf("You ran out of guesses!")
	gameOver(target, totalGuesses, false)
}

// gameOver handles the behavior of the end state when a player wins or loses.
func gameOver(correctNumber int, guessesLeft int, win bool) {
	if win {
		fmt.Printf("You win! You got the right number: %v in %v guesses!\n", correctNumber, guessesLeft+1)
	} else {
		fmt.Printf("\nGame Over! Womp womp! The correct number was %v.\n", correctNumber)
	}

	//Play Again input and validation
	fmt.Printf("\nWould you like to play again? (%s/%s): ", PlayAgainYes, PlayAgainNo)
	var playAgainInput string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		playAgainInput = scanner.Text()
		playAgainInput = strings.ToUpper(playAgainInput) // Make it case-insensitive

		if playAgainInput == PlayAgainYes {
			fmt.Println("The grind never stops")
			return // Return to the main loop to start a new game
		} else if playAgainInput == PlayAgainNo {
			playing = false // Exit the main loop
			fmt.Println("Thank you for playing!! Goodbye!")
			return
		} else {
			fmt.Printf("Invalid input. Please enter '%s' or '%s'.\n", PlayAgainYes, PlayAgainNo)
		}
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
	if !match {
		fmt.Println("Invalid input. Name must contain only letters.")
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
