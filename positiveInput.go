// Created by: haokai
// Created on: May 2021
// This program displays, "Positive-and-Negative"

package main

import (
	"fmt"
)

func main() {

	// This function does addition
	var number int

	// input
	fmt.Println("This program is about Positive and Negative.")
	fmt.Println()
	fmt.Print("Enter integers of number: ")
	fmt.Scanln(&number)

	// detect
	if number < 0 {
		// output of right answer
		fmt.Println("It is negative.")
	} else {
		// output of wrong answer
		fmt.Println("It is positive.")
	}
	fmt.Println("\nDone.")
}