package main

import "fmt"

type human struct {
	sex  int
	name string
}

type human2 struct {
	sex  int
	name string
}

type teacher struct {
	human
	name string
}

type student struct {
	human
	human2
	sex  int
	name string
}

func main() {
	// 内嵌匿名类，默认名称使用和类型一样
	teacher := teacher{name: "teacher", human: human{sex: 0}}
	teacher.sex = 1
	fmt.Println(teacher)

	student := student{name: "student", human: human{sex: 1, name: "human"}}
	student.sex = 1
	fmt.Println(student.name)

	// 内嵌的类型中不能有重复的属性，如A{B,C}, B{name}, C{name}, 这样会报错
	// 如果A下有name，默认取A的name, 并且不会报错了
}
