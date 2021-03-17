package main

import "fmt"

func test(p *int) {
	fmt.Printf("%p,%v\n", &p, p)
}
func main() {
	a := 0x100
	p := &a
	fmt.Printf("%p,%v\n", &p, p)
	test(p)
	//0xc0000b0018,0xc0000ae008
	//0xc0000b0028,0xc0000ae008
	//尽管实参和形参都指向同一目标，但指针传递时依然被复制

}
