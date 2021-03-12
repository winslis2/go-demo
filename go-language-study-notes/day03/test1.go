package main

import "fmt"

func main() {
	x := 2 << 2
	fmt.Println(x)
	a := 0
	a++ //不能++a
	p := &a
	*p++
	fmt.Println(a) //2
	//c := map[string]int{"a":1}
	//fmt.Println(&c["a"]) //cant' take the address of c["a"]

}
