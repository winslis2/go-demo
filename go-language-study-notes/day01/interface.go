package main

import "fmt"

type user1 struct {
	name string
	age  byte
}

//这里的参数是放到func后的，不是放到print后的,当然print也可以跟参数
//func后好像只能跟一个参数
func (u user1) print(s string) {
	fmt.Println(u)
	fmt.Println(s)
}

type Printer interface {
	print(s string)
}

func main() {
	var u user1
	u.name = "lis2"
	u.age = 18

	var p Printer
	p = u
	p.print("ss")
}
