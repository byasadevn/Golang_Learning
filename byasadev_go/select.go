package main

import (
	"fmt"
)

func process1(ch chan string) {
	//time.Sleep(150 * time.Millisecond)
	ch <- "proccess 1"
}
func process2(ch chan string) {
	//time.Sleep(100 * time.Millisecond)
	ch <- "process 2"
}

func select() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process1(ch1)
	go process2(ch2)
	select {
	case p1 := <-ch1:
		fmt.Println(p1)
	case p2 := <-ch2:
		fmt.Println(p2)

	}

}
