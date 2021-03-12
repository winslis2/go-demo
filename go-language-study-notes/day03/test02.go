package main

import "fmt"

func main() {
	a := struct {
		x int
		y int
	}{1, 2}
	a.x = 100
	p := &a
	fmt.Println(p)   //&{100,2}
	p.x += 100       //==a.x+=100
	fmt.Println(p.x) //200
}
