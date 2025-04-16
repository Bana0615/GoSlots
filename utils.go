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

	//Equivilant to a while loop
	for true {
		fmt.Printf("Enter your bet (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Insufficient balance. Try again.")
		} else {
			break
		}
	}

	return bet
}