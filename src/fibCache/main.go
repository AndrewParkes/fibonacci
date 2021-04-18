package main

import (
	"fmt"
	"time"
)

var cache = make(map[int]int)

func main() {
	start := time.Now()

	fmt.Println("fib:", fib(50))

	duration := time.Since(start)
	fmt.Println("Nanoseconds:", duration.Nanoseconds())
	fmt.Println("map:", cache)
}

func fib(n int) int {


	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		if cache[n] != 0 {
			return cache[n]
		}
		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}
}
