package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func (user User) getName() {
	fmt.Println(user.Name)
}

func (user User) GetName(name string, age int) {
	fmt.Printf("hi, %s(age: %v), my name is %s\n", name, age, user.Name)
}

func main() {
	user := User{"zhangsan"}
	userType := reflect.ValueOf(user)

	// 只能获取public的method
	fmt.Println(userType.NumMethod())
	fmt.Println(userType.IsValid())
	fmt.Println(userType.NumField())

	method := userType.MethodByName("GetName")
	params := []reflect.Value{reflect.ValueOf("lisi"), reflect.ValueOf(18)}
	method.Call(params)
}
