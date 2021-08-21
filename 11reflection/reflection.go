package main

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	// 大写就是public， 小写就private
	Name string
	Age  int
}

func main() {
	teacher := Teacher{"wang", 18}

	//info(teacher)
	info2(&teacher)
}

func info2(o interface{}) {
	typpOf := reflect.TypeOf(o)
	if kind := typpOf.Kind(); kind != reflect.Struct {
		fmt.Println(kind)
	}
}

func info(teacher interface{}) {
	// reflect.TypeOf() 函数可以获得任意值的类型对象
	ref := reflect.TypeOf(teacher)
	fmt.Println(ref.Name())

	value := reflect.ValueOf(teacher)
	fmt.Println("Fields: ")

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		val := value.Field(i).Interface()
		fmt.Printf("name = %s, type = %s, value = %v\n", field.Name, field.Type, val)
	}

	for i := 0; i < ref.NumMethod(); i++ {
		m := ref.Method(i)
		fmt.Printf("method name = %s, method type =  %s\n", m.Name, m.Type)
	}
}

func (t Teacher) Hello() {
	fmt.Println("Hello World!")
}
