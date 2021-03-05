package main

import (
	"fmt"
	"time"
)

func main() {
	go task(2)
	go task(5)
	time.Sleep(time.Second * 6)
}

func task(id int) {
	for i := 0; i < 5; i++ {
		//使用格式化输出的时候用\n进行换行，C语言也是同样的
		fmt.Printf("id=>%d:i=>%d\n", id, i)
		time.Sleep(time.Second)
	}
}
