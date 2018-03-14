package main

import (
	"fmt"
	"time"
)

func brewCoffee(c chan<- string, i int) {
	// cashier receive order
	time.Sleep(5 * time.Millisecond)
	coffee := fmt.Sprintf("order: %d", i)

	// barista brew coffee
	time.Sleep(100 * time.Millisecond)
	c <- fmt.Sprintf("%s %s", coffee, "espresso")
}
func serveCoffee(c <-chan string, volumn int) (container []string) {
	for i := 1; i <= volumn; i++ {
		coffee := <-c
		time.Sleep(5 * time.Millisecond)
		container = append(container, fmt.Sprintf("%s %s", coffee, "ready :)"))
	}
	return
}
func order(volumn int) (container []string) {
	c := make(chan string)
	for i := 1; i <= volumn; i++ {
		go brewCoffee(c, i)
	}
	container = serveCoffee(c, volumn)
	return
}
func main() {
	volumn := 200
	start := time.Now()
	container := order(volumn)
	for _, cup := range container {
		fmt.Println(cup)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
