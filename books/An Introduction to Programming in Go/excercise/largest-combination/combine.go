package combine

import (
	//"fmt"
	"strconv"
)

func swap(input []int, i int, j int) {
	temp := input[j]
	input[j] = input[i]
	input[i] = temp
}
func isSwapAtIndex(index int, left string, right string) bool {
	lengthOfLeft := len(left)
	lengthOfRight := len(right)
	if lengthOfLeft > index && lengthOfRight > index {
		return left[index] > right[index]
	}
	if lengthOfRight == 1 && lengthOfLeft > 1 && left[1] > '0' {
		return true
	}
	if lengthOfLeft == 1 && lengthOfRight > 1 && right[1] >= '0' {
		return true
	}
	if lengthOfLeft == 2 && lengthOfRight > 2 && right[1] >= '0' {
		return true
	}
	return false
}
func diffAtIndex(left string, right string) (diffIndex int) {
	endIndex := len(right)
	if len(right) > len(left) {
		endIndex = len(left)
	}
	for i := 0; i < endIndex; i++ {
		if right[i] != left[i] {
			break
		}
		diffIndex++
	}
	return
}
func sortCharactor(input []int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			right := strconv.Itoa(input[i])
			left := strconv.Itoa(input[j])
			startIndex := diffAtIndex(left, right)
			if isSwapAtIndex(startIndex, left, right) {
				swap(input, i, j)
			}
		}
	}
}
func combinedNumber(input ...int) string {
	sortCharactor(input)
	output := ""
	for _, v := range input {
		output += strconv.Itoa(v)
	}
	//fmt.Println(input)
	return output
}
