package main

type woker interface {
	walk()
}
type laser int

func (l *laser) walk() {

}

func shout(t woker) {

}
func main() {

	a := laser(2) //将int类型的2转化为laser类型
	shout(&a)     //应为walk方法接收者是laser指针，此处必须是指针传递，如果walk方法接收者是普通，则可以传指针，也可以传非指针
}
