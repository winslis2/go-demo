package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil { //要进行nil的判断，所以以后在出现声明一个指针的时候要判断是否为nil
		return
	}
	p.age++ //这个地方会进行自动解引用
}

func main() {
	var a *int
	fmt.Println(a) //nil
	//fmt.Println(*a) //打印nil的指针引用会出panic错误

	var nobody *person
	nobody.birthday()

	interfaceNil()
}

func fn() {
	//函数类型的默认值是nil ,必要的时候要提供默认行为，也即是给这个nil赋值
	var fun func(a, b int) int
	fmt.Println(fun == nil) //true
}

func interfaceNil() {
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) //	<nil> <nil> true
	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil) //	*int <nil> false //类型是*int类型，虽然值相同，但是用等号比价还是false
}
