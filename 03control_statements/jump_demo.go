package main

import "fmt"

func main() {
	// goto 跳到 LABEL1 同级位置, 从头执行
	// break 跳到 LABEL1 同级位置，不执行下一个语句
	// continue 跳到 LABEL1 同级位置， 继续上一次执行的位置往下走
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

LABEL2:
	for i := 0; i < 10; i++ {
		for {
			// 一直执行最外层for
			fmt.Println(i)
			continue LABEL2
		}
	}
	fmt.Println("OK2")
}
