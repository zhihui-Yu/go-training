package main

import "fmt"

func main() {
	var m map[int]string = make(map[int]string)
	fmt.Println(m)

	m2 := make(map[int]string)
	m2[1] = "OK"
	delete(m2, 1)
	fmt.Println(m2)

	m3 := make(map[int]map[int]string)
	m3[1] = make(map[int]string)
	m3[1][1] = "OK"
	// ok 代表是否有初始化
	a, ok := m3[2][1]
	if !ok {
		m3[2] = map[int]string{1: "OK"}
	}
	fmt.Println(m3)
	fmt.Println(a, ok)

}
