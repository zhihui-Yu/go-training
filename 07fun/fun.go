package main

import "fmt"

func main() {
	// 一切皆函数
	fmt.Println(A())
	fmt.Println(A1(1))
	fmt.Println(A2(1, 2, "3"))
	fmt.Println(A3(1, 2))
	fmt.Println(A3(1, 2, 3, 4, 5))

	// A 的复制品
	a := A
	fmt.Println(a())

	a2 := func() { fmt.Println("func a") }
	a2()

	// 闭包， 返回一个函数
	a3 := func(x int) func(int) int {
		return func(int) int {
			fmt.Println("func a3")
			return 2
		}
	}
	a3(2)(2)
}

func A() int {
	return 1
}

func A1(a int) int {
	return a
}

func A2(a, b int, c string) (int, int, string) {
	return a, b, c
}

func A3(a ...int) (int, int) {
	return a[0], a[1]
}
