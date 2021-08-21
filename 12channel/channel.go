package main

import "fmt"

func main() {
	// 声明子线程
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO !!!")
		// 给通道写值
		c <- true
	}()
	// 默认是不会关闭的
	close(c)
	// 让主线程阻塞在这里， 等待值得写入
	<-c
}
