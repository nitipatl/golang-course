package main

import "fmt"

// func main() {
// 	fmt.Println(1)
// 	fmt.Println(2)
// 	fmt.Println(3)
// 	fmt.Println(4)
// 	fmt.Println(5)
// 	fmt.Println(6)
// 	fmt.Println(7)
// 	fmt.Println(8)
// 	fmt.Println(9)
// 	fmt.Println(10)
// }

// for loop
func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	for i := 1; i <= 10; i++ {
		// fmt.Print(i)
		// if i%2 == 0 {
		// 	fmt.Println(" Odd")
		// } else {
		// 	fmt.Println(" Even")
		// }
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		default:
			fmt.Println("No number")
		}
	}
}
