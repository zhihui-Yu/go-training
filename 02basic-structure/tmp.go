package main

import (
	"fmt"
	"math"
)

type (
	// 可重写已有的type
	byte int8
	rune int32
	文本   string
)

func main() {
	var temp float32 = 100.1
	fmt.Println(temp)
	// 不能进行转换
	//tempB := bool(temp)
	//fmt.Println(tempB)

	// 赋值
	//效果同下 var ne int = 1
	ne := 2 // 系统自动推断
	fmt.Println(ne)

	// 同时声明多个变量
	//var tmp1, tmp2, tmp3 = 1, 2, 3
	// _ 是一个空白符
	tmp1, _, tmp2, tmp3 := 1, 2, 3, 4
	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)

	var tmp 文本
	tmp = "中文类型名"
	fmt.Println(tmp)

	var a int
	println(a)

	var b bool
	println(b)

	var str string
	println(str)

	var arr [1]int
	fmt.Println(arr)

	var arr2 [1]bool
	fmt.Println(arr2)

	var arr3 [1]byte
	fmt.Println(arr3)

	// 检查溢出
	fmt.Println(math.MinInt32)
}
