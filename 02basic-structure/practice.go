package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 68
	// 程序会将int 值变为 该数字在 ASCII 码中对应的符号
	b := string(a)
	fmt.Println(b)
	// 如果想转为对应的数字，使用如下方法
	c := strconv.Itoa(a)
	// 将string转int
	fmt.Println(c)
	a, _ = strconv.Atoi(c)
	fmt.Println(a)
}
