package main

import "fmt"

func main() {

	// Number
	fmt.Println("1.0 + 1.5 =", 1.0+1.5)
	fmt.Println("1 + 1 =", 1+1)

	// String
	fmt.Println(len("Hello World")) // A space is also considered a character
	fmt.Println("Hello World"[1])   // represent to a byte
	fmt.Println("Hello " + "World") // string will be concat not addition

	// Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
