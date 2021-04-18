package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var cache = sync.Map{}

func main() {
	start := time.Now()

	fmt.Println("fib:", fib(50))

	duration := time.Since(start)
	fmt.Println("Nanoseconds:", duration.Nanoseconds())

	printMap()

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
		result, ok := cache.Load(n)

		if !ok {

			channel1 := make(chan int)
			go fibChannel(n-1, channel1)

			channel2 := make(chan int)
			go fibChannel(n-2, channel2)

			val := <- channel1 + <- channel2

			cache.Store(n, val)
			channel <- val
		} else {
			channel <- result.(int)
		}
	}
}

func printMap() {
	m := map[string]interface{}{}
	cache.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value
		return true
	})

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}