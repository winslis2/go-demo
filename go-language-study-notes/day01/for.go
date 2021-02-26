package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

//go语言中没有while使用for可以实现效果
func while_test() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
