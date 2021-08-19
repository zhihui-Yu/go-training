package main

import "fmt"

func main() {
	a := [10]int{1}
	fmt.Println(a)

	// 指定索引赋值
	b := [10]int{9: 1}
	fmt.Println(b)

	// 不指定长度
	c := [...]int{1, 2, 3, 4}
	fmt.Println(c)

	d := [...]int{99: 1}
	var p = &d
	fmt.Println(d)
	fmt.Println(p)

	// 存储指针
	x, y := 1, 2
	e := [...]*int{&x, &y}
	fmt.Println(e)

	// new 创建的是 指向数组的指针， 但可以通过索引直接操作
	p1 := new([10]int)
	p1[2] = 2
	fmt.Println(p1)

	// 多维数组
	arr := [2][3]int{
		//{1,1,1},
		{1: 1, 2: 2},
		{2, 2, 2},
	}
	fmt.Println(arr)
}
