package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 1
	of := reflect.ValueOf(&x)
	// Elem 获取指针指向得元素
	if of.Elem().CanSet() {
		of.Elem().SetInt(100)
		fmt.Println("changed")
	}
	fmt.Println(x)
}
