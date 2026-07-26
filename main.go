package main

import (
	"fmt"
	"math/rand"
)

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

func generateSymbolArray(symbols map[string]uint) []string {
	res := []string{}
	for symbol, count := range symbols {
		for i:=0;i<int(count);i++{
			res = append(res, symbol)
		}
	}
	return res
}

func getRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max - min + 1) + min
	return randomNumber
}

func getSpin(reel []string, rows int, cols int) [][]string {
	res := [][]string{}

	for row:=0;row<rows;row++ {
		res = append(res, []string{})
	}

	for col:=0;col<cols;col++ {
		selected := map[int]bool{}
		
		for row:=0;row<rows;row++ {
			for true {
				randomIndex := getRandomNumber(0, len(reel) - 1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					res[row] = append(res[row], reel[randomIndex])
					break
				}
			}
		}
	}
	
	return res 
}

func displaySpin(slot [][]string) {
	fmt.Print("\n ------ THE SLOT ------ \n\n")
	for _, row := range slot {
		fmt.Print("      | ")
		for _, ele := range row {
			fmt.Print(ele, " | ")
		}
		newLine()
	}
}

func main()  {
	var balance uint = 200
	var bet uint
	symbols := map[string]uint {
		"A" : 4,
		"B" : 7,
		"C" : 12,
		"D" : 20,
	}

	// multiplier := map[string]uint {
	// 	"A" : 20,
	// 	"B" : 10,
	// 	"C" : 5,
	// 	"D" : 2,
	// }

	res := generateSymbolArray(symbols)

	getName()
	fmt.Printf("(Starting balance = %d)\n", balance)

	for balance > 0 {
		bet = getBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
		slot := getSpin(res, 3, 3) 
		displaySpin(slot)
	}
	newLine()
	fmt.Printf("You left with Rs %d.\n", balance)
}