package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	age := "18x"
	ageint, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("age not a int")
	} else {
		fmt.Println("age int => ", ageint)
	}

	if match, _ := regexp.MatchString("^[0-9]+$", age); !match {
		fmt.Println("age not a int by regex")
	}

	now := time.Now()
	time := time.Date(2022, time.November, 10, 23, 0, 0, 0, time.Local)
	fmt.Println("go time: ", time.Local())
	fmt.Println("go time: ", now)
}
