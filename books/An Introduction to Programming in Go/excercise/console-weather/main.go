package main

import (
	"fmt"
	"strconv"
)

func createNumberText(number int) (out string) {
	numberS := strconv.Itoa(number)
	numberText := [][3]string{
		{
			" _ ",
			"| |",
			"|_|",
		},
		{
			"   ",
			"  |",
			"  |",
		},
		{
			" _ ",
			" _|",
			"|_ ",
		},
		{
			" _ ",
			" _|",
			" _|",
		},
		{
			"   ",
			"|_|",
			"  |",
		},
		{
			" _ ",
			"|_ ",
			" _|",
		},
		{
			" _ ",
			"|_ ",
			"|_|",
		},
		{
			" _ ",
			"  | ",
			"  |",
		},
		{
			" _ ",
			"|_|",
			"|_|",
		},
		{
			" _ ",
			"|_|",
			" _|",
		},
	}
	for i := 0; i < 3; i++ {
		out += "\n"
		for _, v := range numberS {
			index, _ := strconv.Atoi(string(v))
			out += numberText[index][i]
		}
	}
	out += " C \n"
	return
}
func weatherC(c int, text string) string {
	return createNumberText(c) + text
}
func main() {
	fmt.Println(weatherC(26, "Bangkok few cloud"))
}
