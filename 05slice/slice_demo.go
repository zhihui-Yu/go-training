package main

import "fmt"

func main() {
	// slice 变长数组的解决方案
	// 改变slice的值，会直接影响到底层数组的值
	// 但是如果用append函数且超过原数组最大值时，新的slice指向的是新的地址空间，修改了也不会影响到原有的数组
	s1 := make([]int, 3, 10)
	fmt.Println(s1, cap(s1))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	sa := a[2:5]
	fmt.Println(string(sa))
	fmt.Println(string(a[:3]))
	fmt.Println(string(a[2:3]))

	// 超出 内存分配的最大值，就会重新分配一个内存去放全部值
	s2 := make([]int, 3, 6)
	fmt.Printf("%p", s1)
	s2 = append(s1, 2)
	fmt.Println(s2)

	// copy 函数
	s3 := []int{1, 2, 3, 4, 5, 6, 7}
	s4 := []int{19, 20}
	//copy(s3,s4)
	copy(s3[3:4], s4[1:2]) // 指定copy位置
	fmt.Println(s3)
}
