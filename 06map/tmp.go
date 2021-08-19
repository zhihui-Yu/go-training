package main

import (
	"fmt"
	"sort"
)

func main() {

	sm := make([]map[int]string, 5)
	// value 只是值得复制，不影响原值
	for index, value := range sm {
		value = make(map[int]string, 1)
		value[1] = "OK"
		fmt.Println(index, value)
	}
	// 可以使用 sm[index] 直接操作原数组
	fmt.Println(sm)

	// map sort
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

}
