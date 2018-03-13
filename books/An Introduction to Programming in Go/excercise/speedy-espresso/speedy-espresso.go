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
	coffee = fmt.Sprintf("%s %s", coffee, "espresso")
	c <- coffee
}
func serveCoffee(c <-chan string, volumn int, finised chan<- bool) {
	for i := 1; i <= volumn; i++ {
		coffee := <-c
		time.Sleep(5 * time.Millisecond)
		fmt.Println(coffee, "ready :)")
	}
	finised <- true
}
func order(volumn int, finised chan bool) {
	c := make(chan string)
	go serveCoffee(c, volumn, finised)
	for i := 1; i <= volumn; i++ {
		go brewCoffee(c, i)
	}
}
func main() {
	finised := make(chan bool)
	volumn := 10
	start := time.Now()
	order(volumn, finised)
	<-finised
	end := time.Now()
	fmt.Println(end.Sub(start))
}
