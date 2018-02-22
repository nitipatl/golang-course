package main

import "fmt"

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(xs))
// }

// func average(numbers []float64) float64 {
// 	total := 0.0
// 	for _, v := range numbers {
// 		total += v
// 	}
// 	return total / float64(len(numbers))
// }
func main() {
	x, y := f2()
	fmt.Println(x, y)
}
func f2() (int, int) {
	return 5, 6
}
