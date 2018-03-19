package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan bool) {
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Duration(10 * rand.Float64() * float64(time.Microsecond)))
		ch <- true
	}
	close(ch)
}

func consumer(ch chan bool) {
	for _ = range ch {
		time.Sleep(time.Duration(10 * rand.Float64() * float64(time.Microsecond)))
	}
}

func unbuffered() {
	ch := make(chan bool, 0)
	go producer(ch)
	consumer(ch)
}

func buffered() {
	ch := make(chan bool, 1)
	go producer(ch)
	consumer(ch)
}

func main() {
	startTime := time.Now()
	unbuffered()
	fmt.Printf("unbuffered: %vn", time.Since(startTime))
	startTime = time.Now()
	buffered()
	fmt.Printf("buffered: %vn", time.Since(startTime))
}
