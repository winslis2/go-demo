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
	//定义struct可以写name,也可以按顺序不写
	book1 := Book{"PHP是世界上最好的语言", "lis2", "PHP", 1}
	//+v和 v的区别
	fmt.Printf("%v\n", book1)  //{PHP是世界上最好的语言 lis2 PHP 1}
	fmt.Printf("%+v\n", book1) //{title:PHP是世界上最好的语言 auth:lis2 subject:PHP book_id:1}&{hshhs lis2 PHP 1}

	//fmt.Println(book1.title)
	//fmt.Println(book1)
	struct_pointer_test(&book1)
	fmt.Println(book1)
}

//可以传指针，也可以不传指针 ,传递指针会修改原来的值
func struct_pointer_test(book *Book) {
	book.title = "hshhs"
	//结构体指针可以直接使用
	fmt.Println(book)
}
