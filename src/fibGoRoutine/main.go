package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("fib:", fib(50))

	duration := time.Since(start)
	fmt.Println("Nanoseconds:", duration.Nanoseconds())
}

func fib(n int) int {

	channel:= make(chan int)
	go fibChannel(n, channel)
	val :=	<-channel

	return  val
}

func fibChannel(n int, channel chan int){

	if n == 1 {
		channel <- 1
	} else if n == 0 {
		channel <- 0
	} else {

		channel1 := make(chan int)
		channel2 := make(chan int)

		go fibChannel(n-1, channel1)
		go fibChannel(n-2, channel2)

		val := <- channel1 + <- channel2

		channel <- val
	}
}