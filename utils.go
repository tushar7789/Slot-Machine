package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Print("\n\nWelcome to Jim's Casino...\n\n")
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)

	if err != nil {
		return ""
	}

	newLine()
	fmt.Printf("Welcome %s, let's play!\n\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Printf("Enter your bet amount or 0 to quit (balance : Rs %d) : ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be more than the Balance!")
		}else  {
			break
		}
	}

	return bet
}
