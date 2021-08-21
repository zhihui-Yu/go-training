package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := person{
		name: "joe",
		age:  10,
	}
	a.name = "joe"
	a.age = 18
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	A2(&a)
	fmt.Println(a)

}

// A 值传递， 内部修改不影响外部
func A(per person) {
	per.name = "A"
	per.age = 20
}

// A2 指针传递，内部修改影响外部
func A2(per *person) {
	per.name = "A"
	per.age = 20
}
