package main

import "fmt"

func main() {

	fmt.Println("Welcome to my quiz game!")

	println("Enter you name here : ")

	var name string
	fmt.Scan(&name)
	// asks for user input

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: \n ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You can't play")
	}

	score := 0
	numQuestions := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090\n")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		score++
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have?\n")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		score++
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v\n", score, numQuestions)
	percent := (float64(score) / float64(numQuestions)) * 100
	fmt.Printf("You scored: %v%%\n", percent)

}
