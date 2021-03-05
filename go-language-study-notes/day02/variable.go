package main

import "fmt"

var x1 = 100

func main() {
	//变量赋值的时候可以给不同的类型赋值
	var s, i = "lis2", 11

	//使用多行赋值
	var (
		x, y   int
		s1, i1 = "lis2", 11
	)
	fmt.Println(s, i, x, y, s1, i1)

	fmt.Println(&x1, x1)
	//使用等号是修改
	x1 = 101
	fmt.Println(&x1, x1)
	//使用:=是从新定义同名局部变量
	x1 := 102
	fmt.Println(&x1, x1)

	fmt.Println("退化赋值---------")
	degenerate()

}

//:=退化赋值  多个error也有这样的效果
func degenerate() {
	i := 100
	fmt.Println(&i, i)
	//这个时候i的:=退化为赋值
	i, y := 1, 2
	fmt.Println(&i, i)
	fmt.Println(y)
	{
		//不是同一个作用域不会退化，会生成一个局部变量
		i, y1 := 1, 2
		fmt.Println(&i, i, y1)
	}
}
