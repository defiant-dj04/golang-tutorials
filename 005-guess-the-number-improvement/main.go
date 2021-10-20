package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// value assigned to a constant can never change
// constant don't have to be used in our program, for variables if they remain unused compiler will throw an error
const prompt = "and don't type your number in, just press ENTER when ready."

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	// create our reader variable, which reads input from standard input (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display the welcome instructions
	fmt.Println("Guess the Number Game:")
	fmt.Println("----------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}

func main() {
	// seed the random number generator to produce new and unique random numbers everytime we execute our program
	rand.Seed(time.Now().UnixNano())

	// rand generates a number between 0 an whatever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and
	// no greater than 10 (multiplying by 1 makes the game a bit silly)
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber * secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}
