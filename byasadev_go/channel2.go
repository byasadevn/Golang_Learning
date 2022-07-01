package main

import "fmt"

func receiveData(data chan int) {
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)
}

func channel2() {
	ch := make(chan int)
	go receiveData(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Printf("The ok: %v and value: %v\n", ok, v)
	}
}
