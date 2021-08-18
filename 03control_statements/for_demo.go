package main

import "fmt"

func main() {
	// 无限循环
	//for {
	//	fmt.Println("a")
	//}

	// 可以在for上赋值
	for a := 1; a < 10; {
		fmt.Println(a)
		a = a + 1
	}
}
