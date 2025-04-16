package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
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

func getBet(balance uint) uint {
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

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}

	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, symbol)
		}
	}

	return symbolArray
}

func getSpin(reel []string, rows int, cols int) [][]string {
	spin := [][]string{}

	//Create empty slice
	for i := 0; i < rows; i++ {
		spin = append(spin, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}

		for row := 0; row < rows; row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					spin[row] = append(spin[row], reel[randomIndex])
					break
				}
			}

			spin[col] = append(spin[col], reel[row])
		}
	}

	return spin
}

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for i, symbol := range row {
			fmt.Printf("%s ", symbol)
			if i != len(row)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// Setting odds for symbols to show up
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	// Bet multiplier if you get a line of the symbols
	// multipliers := map[string]uint{
	// 	"A": 20,
	// 	"B": 10,
	// 	"C": 5,
	// 	"D": 2,
	// }
	symbolArray := generateSymbolArray(symbols)
	spin := getSpin(symbolArray, 3, 3)
	printSpin(spin)

	balance := uint(200)

	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
	}

	fmt.Printf("You left with, %d.\n", balance)
}
