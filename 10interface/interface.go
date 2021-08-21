package main

import (
	"fmt"
)

type human interface {
	get_name() string
}

type teacher struct {
	name string
}

func main() {
	// go 中没有明显的继承关系，只要谁实现了接口的方法，就可以认为他实现了接口
	teacher := teacher{name: "joe"}
	fmt.Println(teacher.get_name())
	convert(teacher)
}

func (teacher *teacher) get_name() string {
	fmt.Println("name -> ", teacher.name)
	return teacher.name
}

func convert(human interface{}) {
	// 判断是否某个具体实现类
	//if man, ok := human.(teacher); ok {
	//	fmt.Println("convert: ", man.name)
	//	return
	//}
	//fmt.Println("unknown")

	// 让编译器去做判断
	switch v := human.(type) {
	case teacher:
		fmt.Println("convert: ", v.name)
	default:
		fmt.Println("unknown")
	}
}
