package main

import "fmt"

type person2 struct {
	age int
}

func (p *person2) birthday() {
	//下面的用法是自动解引用
	p.age++
}
func main() {
	b := person{1}
	b.birthday()

	//使用指针作为接收者的策略应始终如一，如果一中类型的某些方法需要使用指针作为接收者，就应该为这种类型的所有方法都设置为指针接收者

	a := []int{1, 2, 3}
	test4(&a)
	fmt.Println(a)
}

func test1(m map[string]string) {

}

//也会直接修改原值
func test2(s []int) {
	s[0] = 2
}

func test3(s []int) {
	s = s[0:1] //不会改变外部的值
}

func test4(s *[]int) {
	*s = (*s)[0:1] //会改变外部的值
}
