package main

import "fmt"

type IN int

func main() {
	var i IN
	i.Increase(100)
	fmt.Println(i)
}

func (i *IN) Increase(x int) {
	// 由于int 和 IN 是同类型，需要转成IN才能做运算
	*i += IN(x)
}
