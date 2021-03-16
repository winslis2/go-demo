package main

import "fmt"

func main() {
	//匿名函数的使用
	func(s string) {
		fmt.Println(s)
	}("hello  world")
}
