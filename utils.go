package main

import (
	"fmt"
)

func GetName() string {
	name := ""

	fmt.Println("Welcome to Banas Casino...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	fmt.Printf("Hello %s, lets play!\n", name)

	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for {
		// Prompt the user, including current balance
		fmt.Printf("Enter your bet (balance = $%d, enter 0 to quit): ", balance)

		// Attempt to scan directly into the uint variable
		_, err := fmt.Scan(&bet)

		if err != nil {
			// If scanning fails (e.g., user typed text), print an error.
			fmt.Println("Invalid bet. Please enter a whole non-negative number.")
			// Continue the loop to ask for input again
			continue
		}

		// If Scan succeeded, check if the bet is within the player's balance
		if bet > balance {
			fmt.Println("Insufficient balance. Try again.")
		} else {
			// Bet is a valid uint and within balance (or 0 to quit), exit the loop
			break
		}
	}

	return bet
}
