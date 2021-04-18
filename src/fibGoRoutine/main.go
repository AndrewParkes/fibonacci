package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(10)
	start := time.Now()

	fmt.Println("fib:", fib(5))

	duration := time.Since(start)
	fmt.Println("Nanoseconds:", duration.Nanoseconds())
}

func fib(n int) int {

	channel:= make(chan int)
	go fibCo(n, channel)
	val :=	<-channel

	return  val
}

func fibCo(n int, channel chan int){

	if n == 1 {
		channel <- 1
	} else if n == 0 {
		channel <- 0
	} else {

		channel1 := make(chan int)
		channel2 := make(chan int)

		go fibCo(n-1, channel1)
		go fibCo(n-2, channel2)

		val := <- channel1 + <- channel2

		channel <- val
	}
}