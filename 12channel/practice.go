package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	time.Sleep(2 * time.Second)
	c := make(chan string)
	go func() {
		for {
			fmt.Println(<-c)
			c <- fmt.Sprintf("Go here")
		}
	}()

	for i := 0; i < 3; i++ {
		c <- fmt.Sprintf("Main here")
		fmt.Println(<-c)
	}
	close(c)
}
