package main

import "fmt"

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

func main() {
	getName()
}