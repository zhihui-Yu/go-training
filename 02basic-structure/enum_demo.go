package main

import "fmt"

// iota 每定义一个常量就+1， 从 0 开始， 跟定义顺序没有关系
const (
	a1 = 'A'
	b1
	c1 = iota
	d1
)

const (
	e = iota
)

func main() {
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(d1)
	fmt.Println(e)
}
