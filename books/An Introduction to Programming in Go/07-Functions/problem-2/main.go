package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
func half(number int) (int, bool) {
	return number / 2, number%2 == 0
}
