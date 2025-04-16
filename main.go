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

func getRandomNumber(min int, max int) int {
	return rand.Intn(max-min+1) + min
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
		}
	}

	return spin
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

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true

		checkSymbol := row[0]

		//Skip the first element in the row
		for _, symbol := range row[1:] {
			if symbol != checkSymbol {
				win = false
				break
			}
		}

		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}

	return lines
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
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	symbolArray := generateSymbolArray(symbols)

	balance := uint(200)

	getName()
	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := getSpin(symbolArray, 3, 3)
		printSpin(spin)

		//Check if they win and update balance
		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win

			if multi > 0 {
				fmt.Printf("You won $%d, (%dx) on line %d\n", win, multi, i+1)
			}
		}
	}

	fmt.Printf("You left with, %d.\n", balance)
}
