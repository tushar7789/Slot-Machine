package main

import "fmt"

func getName() string {
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

func getBet(balance uint) uint {
	var bet uint

	for true {
		fmt.Print("Enter your bet amount or 0 to quit : ")
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be more than the Balance!")
		}else  {
			break
		}
	}

	return bet
}

func newLine() {
	fmt.Print("\n")
}



func main()  {
	var balance uint = 200
	var bet uint

	getName()

	fmt.Printf("(Starting balance = %d)\n", balance)

	for balance > 0 {
		bet = getBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
	}

	newLine()
	fmt.Printf("You left with Rs%d.\n", balance)
}