package main

import "fmt"

func main() {
	AX()
	B()
	C()
}

func AX() {
	fmt.Println("Func A")
}

func B() {
	// 这步类似try catch， 判断出错了然后去recover()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
			fmt.Println(err)
		}
	}()
	// 让程序报异常
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
