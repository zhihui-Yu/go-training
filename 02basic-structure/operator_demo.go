package main

import "fmt"

func main() {
	// ^ 一元运算符
	fmt.Println(^1)
	// ^ 二元运算符
	fmt.Println(2 ^ 1)

	fmt.Println(!false)

	fmt.Println(1 << 10 << 10 >> 10)

	fmt.Println(1 | 2)

	fmt.Println(2 & 5)

	// 6 -> 0110
	// 11   1011
	//      0100  第二行是1，就要变成0， 第二行是 0, 就看第一行的数
	fmt.Println(6 &^ 11)
}
