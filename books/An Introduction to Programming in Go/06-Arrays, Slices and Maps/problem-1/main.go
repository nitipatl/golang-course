package main

import "fmt"

func main() {
	// x := [5]int64{1, 2, 3, 4, 5}
	// x := make([]int, 3, 100) // [0, 0, 0, x, ]

	// fmt.Println(x)
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}
