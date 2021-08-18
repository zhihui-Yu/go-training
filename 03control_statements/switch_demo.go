package main

import "fmt"

func main() {
	a := 1

	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	// 默认执行完一个就不执行下面的了， 但可以用关键字fallthrough
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}

	// 声明变量
	switch b := 0; {
	case b >= 0:
		fmt.Println("b>=0")
		fallthrough
	case b >= 1:
		fmt.Println("b>=1")
	default:
		fmt.Println("None")
	}
}
