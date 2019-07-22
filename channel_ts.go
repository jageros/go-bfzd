package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("==================")
		close(ch)
	}()
	//time.Sleep(time.Second * 3)
	n, ok := <-ch
	if ok {
		fmt.Println(n)
	} else {
		fmt.Printf("error:%d", n)
	}
}
