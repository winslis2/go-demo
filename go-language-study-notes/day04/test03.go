package main

import "fmt"

func main() {
	//匿名函数的使用
	func(s string) {
		fmt.Println(s)
	}("hello  world")

	add := test2() //返回的是一个函数
	add(1, 2)
}

func test2() func(int, int) int {
	return func(i1 int, i2 int) int {
		return i1 + i2
	}
}
