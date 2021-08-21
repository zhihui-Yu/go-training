package main

import "fmt"

type customer struct {
	name    string
	age     int
	contact struct {
		phone, city string
	}
}

func main() {
	a := &struct {
		name string
		age  int
	}{name: "name", age: 20}
	// 可以不用 &a.name = xxx
	a.name = "changed"
	fmt.Println(a)

	// 匿名结构初始化
	a2 := customer{name: "customer", age: 18, contact: struct{ phone, city string }{phone: "xxxx", city: "xxxxx"}}
	a2.contact.phone = "phone"
	a2.contact.city = "city"
	fmt.Println(a2)
}
