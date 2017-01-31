package main

import (
	"time"
	"fmt"
)
func fibonacciThread(n int, c chan int) {
	if n < 2 {
		c <- n
	} else {
		c_rec := make(chan int)
		go fibonacciThread(n-1, c_rec)
		go fibonacciThread(n-2, c_rec)
		r1, r2 := <-c_rec, <-c_rec
		c <- r1 + r2
	}
}

func Run(n int) int {
	c := make(chan int)
	go fibonacciThread(n, c)
	return <-c
}

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println(Run(i));
		time.Sleep(2 * time.Second)
	}
}