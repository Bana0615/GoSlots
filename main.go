package main

import (
	"fmt"
)

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
	symbolArray := GenerateSymbolArray(symbols)

	balance := uint(200)

	GetName()
	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := GetSpin(symbolArray, 3, 3)
		PrintSpin(spin)

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
