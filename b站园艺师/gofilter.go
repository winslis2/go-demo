package main

import (
	"fmt"
	"strings"
)

//chan写左边是要往chan中写值，chan在右边说明是要从chan中取值
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream) //关闭的是c0
	//downstream <- ""
}

func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream) //关闭的是c1
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func filterGopher1(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printerGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher1(c0, c1)
	printerGopher(c1)
}
