// 当前程序的包名
package main

import (
	"fmt"
)

// import std fmt -> std 别名，可以用 std.printf() 代替 fmt.Printf()

// 外部可以使用的方法或者变量肯定都是首字母大写的，如fmt的Println

// 常量定义
const (
	PI   = 3.14
	NAME = "xxxx"
)

// 全局变量的声明与赋值 (文件内可用)
var (
	name = "simple"
	call = "call"
)

// 一般类型声明
type (
	newType int
	type2   string
)

// 结构的声明
type goHer struct {
}

// 接口声明
type golang interface {
}

// 由 main 函数作为程序入口点启动
func main() {
	fmt.Printf("Hello World!")
	println("你好，世界！")
}
