package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher2(i, c)
	}

	timeout := time.After(2 * time.Second)
	//一个通道等待多个goroutine
	for i := 0; i < 5; i++ {
		//select从多个通道中取值
		select {
		case gopherID := <-c:
			fmt.Println("gopher", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience run out")
			return
		}
	}
}

func sleepyGopher2(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	//fmt.Println("--snore--",i)
	c <- i
}
