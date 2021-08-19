package main

import "fmt"

func main() {
	a := [...]int{3, 1, 2, 4, 6}
	fmt.Println(a)
	var len = len(a)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)
}
