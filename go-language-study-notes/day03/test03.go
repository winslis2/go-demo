package main

import "fmt"

func main() {
	type data struct {
		x int
		s string
	}

	var a data = data{1, "name"} //等号后面要写data的
	b := data{
		1,
		"name", //花括号换行的时候必须以逗号结尾
	}

	fmt.Println(a, b)
}
