package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Teacher
}

type Teacher struct {
	// 大写就是public， 小写就private
	Name string
	Age  int
}

func main() {
	user := User{"zhangsan", Teacher{"wangwu", 18}}
	fmt.Println(user)
	// 获取user的类型
	kind := reflect.TypeOf(user).Kind()
	fmt.Println(kind)

	// 通过下标获取子类中的属性
	fmt.Println(reflect.TypeOf(user).FieldByIndex([]int{1, 1}))
}
