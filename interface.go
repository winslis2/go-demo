package main

import "fmt"

type Phone interface {
	call()
}

type Iphone struct {
}

//声明属于Iphone这个struct
func (i Iphone) call() {
	fmt.Println("implement me")
}

func call(phone Iphone) {
	fmt.Println("i'm iphone")
}

func main() {
	var phone Phone
	phone = new(Iphone)
	phone.call()
}
