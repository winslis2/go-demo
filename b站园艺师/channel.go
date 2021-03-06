package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	//一个通道等待多个goroutine
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher", gopherID, "has finished sleeping")
	}
}

func sleepyGopher(i int, c chan int) {
	time.Sleep(3 * time.Second)
	//fmt.Println("--snore--",i)
	c <- i
}
