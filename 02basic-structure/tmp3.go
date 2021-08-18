package main

import "fmt"

func main() {
	a := 1
	// *int 声明一个指针，&是取a的地址
	var p *int = &a
	fmt.Println(*p)
}
