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

	if n == 1 {
		return 1
	} else if n == 0 {
		return 0
	} else {
		return fib(n-1) + fib(n-2)
	}
}
