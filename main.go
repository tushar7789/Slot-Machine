package main

import (
	"fmt"
)

func newLine() {
	fmt.Print("\n")
}

func calculateProfit(slot [][]string, multiplier map[string]uint, bet uint) uint {
	var totProfit uint = 0
	var profit uint = 0
	flag := false
	
	for i, row := range slot {
		profit = 0
		flag = true
		for j, symbol := range row {
			if j < len(row)-1 && symbol != row[j+1] {
				flag = false
				break
			}
		}
		if flag {
			profit = profit + multiplier[row[0]]*bet
			fmt.Printf("You made Rs %d (%dx) on line %d!\n", profit, multiplier[row[0]], i+1)
			totProfit += profit
		}
	}

	return totProfit
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

	multiplier := map[string]uint {
		"A" : 20,
		"B" : 10,
		"C" : 5,
		"D" : 2,
	}

	res := GenerateSymbolArray(symbols)
	
	GetName()
	for balance > 0 {
		bet = GetBet(balance)
		if bet == 0 {
			break
		}
		balance = balance - bet
		slot := GetSpin(res, 3, 3)
		DisplaySpin(slot) 
		balance = balance + calculateProfit(slot, multiplier, bet) 
	}
	newLine()
	fmt.Printf("You left with Rs %d.\n", balance)
}