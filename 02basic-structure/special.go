package main

// for to distinguish & *

// &是“取地址运算符”，是从一个变量获取地址
// *是“解引用运算符”，可以简单理解为“从地址取值”， 是&的逆运算
// && 取指针的指针地址

import "fmt"

type Test struct {
	name string
}

func main() {
	testa := Test{"testa"}
	fmt.Println(testa)
	//结果{test}

	testb := &Test{"testb"}
	fmt.Println(testb)
	//结果 &{test}

	testc := &Test{"testc"}
	fmt.Println(*testc)
	//结果 {test}

	testd := &Test{"testd"}
	fmt.Println(&testd)
	//结果 0xc000006030

	fmt.Println(**&*&testd)
	//结果 {testd}

	var a int = 1
	fmt.Println(a)
	//结果 1
	fmt.Println(&a)
	//结果 0xc00000c0d8
}
