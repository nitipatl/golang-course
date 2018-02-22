package main

import "fmt"

func main() {
	fmt.Println(max(1, 2, 30, 4, 5))
}
func max(numbers ...int) (max int) {
	max = numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return
}
