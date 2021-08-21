package main

import "fmt"

type A struct {
	name string
}

type B struct {
	name string
}
type TX int

func main() {
	// 通过给方法绑定类型的方式给类型添加方法
	a := A{}
	a.Print()

	b := B{}
	b.Print()

	var tx TX
	tx.Print()
}

func (b *TX) Print() {
	fmt.Println("TX = ", b)
}

func (a A) Print() {
	a.name = "joe"
	fmt.Println("A = ", a)
}

func (b *B) Print() {
	b.name = "joe"
	fmt.Println("B = ", b)
}
