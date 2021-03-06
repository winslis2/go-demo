package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher1()          //分线路
	time.Sleep(4 * time.Second) //主线路  主程序结束所有的分支也结束
}

func sleepyGopher1() {
	time.Sleep(3 * time.Second)
	fmt.Println("--snore--")
}
