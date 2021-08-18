package main

import "fmt"

const (
	a = 1
	b
	c
)

func main() {
	// var 是全局变量，加载时间是在const之后
	// b,c 默认用了 a 的值
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
