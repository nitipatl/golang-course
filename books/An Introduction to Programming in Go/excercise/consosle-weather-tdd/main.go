package main

import (
	"strconv"
)

func celciusSegmentIsLastLine(line int) bool {
	return line == 2
}

func getNumber(number int) [3]string {
	numbers := [][3]string{
		{
			"",
			"",
			"",
		},
		{
			"",
			"",
			"",
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
	}
	return numbers[number]
}

func celciusSegment(celcius int) string {

	numberS := strconv.Itoa(celcius)
	a, _ := strconv.Atoi(string(numberS[0]))
	b, _ := strconv.Atoi(string(numberS[1]))
	output := "\n" + numberAppend(getNumber(a), getNumber(b))
	return output
}

func weatherCelcius(celcius int, text string) string {
	out := celciusSegment(celcius)
	out += text + "\n"
	return out
}
func numberAppend(a [3]string, b [3]string) (out string) {
	for i := 0; i < 3; i++ {
		out += a[i] + b[i]
		if celciusSegmentIsLastLine(i) {
			out += " c"
		}
		out += "\n"
	}
	return
}

func main() {

}
