package main

import "fmt"
import "math/rand"

func GenerateSymbolArray(symbols map[string]uint) []string {
	res := []string{}
	for symbol, count := range symbols {
		for i:=0;i<int(count);i++{
			res = append(res, symbol)
		}
	}
	return res
}

func GetRandomNumber(min int, max int) int {
	randomNumber := rand.Intn(max - min + 1) + min
	return randomNumber
}

func GetSpin(reel []string, rows int, cols int) [][]string {
	res := [][]string{}

	for row:=0;row<rows;row++ {
		res = append(res, []string{})
	}

	for col:=0;col<cols;col++ {
		selected := map[int]bool{}
		
		for row:=0;row<rows;row++ {
			for true {
				randomIndex := GetRandomNumber(0, len(reel) - 1)
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

func DisplaySpin(slot [][]string) {
	for _, row := range slot {
		for j, ele := range row {
			fmt.Print(ele)
			if j < len(row) - 1{
				fmt.Print(" | ")
			}
		}
		newLine()
	}
}


