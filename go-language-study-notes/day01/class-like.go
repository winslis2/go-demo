package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manage struct {
	//这个地方间接的实现了继承
	user
	title string
}

func main() {
	var m1 manage
	m1.name = "lis2"
	m1.age = 18
	m1.title = "超级管理员"
	fmt.Println(m1)

}
