package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan time.Duration)
	go loop(channel)
	fmt.Println("Arround Here")

	msg := <-channel
	fmt.Println(msg)
}

func loop(channel chan time.Duration) {
	start := time.Now()
	for i := 0; i < 1000000000000; i++ {

	}

	final := time.Now()
	channel <- final.Sub(start)
}
