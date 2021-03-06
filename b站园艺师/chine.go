package main

func main() {
	c := make(chan int)
	//只有在goroutine中加如通道数据才可以
	go test(c)
	<-c
}
func test(c chan int) {
	c <- 1
}
