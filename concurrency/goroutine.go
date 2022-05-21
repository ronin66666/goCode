package main

import (
	"fmt"
	"time"
)

func main() {
	waitTest()
}

//1. wait: 进程退出时不会等待并发任务结束，可用管道（channel）阻塞，然后发出退出信号。
func waitTest() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine done.")
		close(ch)
	}()

	fmt.Println("main ...")
	 <- ch //阻塞 直到读取到数据，或者通道关闭
	fmt.Println("main exit")
}
