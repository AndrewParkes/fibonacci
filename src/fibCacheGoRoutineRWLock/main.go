package main

import (
	"fmt"
	"sync"
	"time"
)

var cache = map[int]int{}

var lock = sync.RWMutex{}

func main() {
	start := time.Now()

	fmt.Println("fib:", fib(20))

	duration := time.Since(start)
	fmt.Println("Nanoseconds:", duration.Nanoseconds())

	fmt.Println(cache)

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

		if read(n) <= 0 {

			channel1 := make(chan int)
			go fibChannel(n-1, channel1)

			channel2 := make(chan int)
			go fibChannel(n-2, channel2)

			val := <- channel1 + <- channel2

			write(n, val)
			channel <- val
		}
			channel <- read(n)
	}
}

func read(n int) int {
	lock.RLock()

	val := cache[n]

	lock.RUnlock()
	return val
}

func write(val int, n int) {
	lock.Lock()

	cache[n] = val

	lock.Unlock()
}

