package main

import "fmt"

type Book struct {
	title   string
	auth    string
	subject string
	book_id int
}

func main() {
	//struct_test()
	struct_test()
}

func struct_test() {
	book1 := Book{"PHP是世界上最好的语言", "lis2", "PHP", 1}
	//fmt.Println(book1.title)
	//fmt.Println(book1)
	struct_pointer_test(&book1)
}

func struct_pointer_test(book *Book) {
	//结构体指针可以直接使用
	fmt.Println(book)
}
