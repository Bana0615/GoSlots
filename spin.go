package main

import (
	"fmt"
	"math/rand"
)

func GenerateSymbolArray(symbols map[string]uint) []string {
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

func GetSpin(reel []string, rows int, cols int) [][]string {
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

func PrintSpin(spin [][]string) {
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
