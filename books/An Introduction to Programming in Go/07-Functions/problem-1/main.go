package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
}
func sum(numbers ...int) (total int) {
	total = 0
	for _, v := range numbers {
		total += v
	}
	return
}
