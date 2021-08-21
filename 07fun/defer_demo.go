package main

import "fmt"

func main() {
	// defer 后调用的先执行
	//fmt.Println("0")
	//defer fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")

	// 都是3， 原因是 函数中的i是引用局部变量的值，在defer执行时，i已经等于3，所以每个语句的i都是3了
	for i := 0; i < 3; i++ {
		fmt.Println("before i = ", i)
		defer func() {
			fmt.Println(i)
		}()
		defer fmt.Println("i = ", i)
		fmt.Println("out i = ", i)
	}
}
