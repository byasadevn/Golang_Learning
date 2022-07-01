package main

import (
	"fmt"
	"sync"
)

var x123 = 0

func incrValue(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x123 = x123 + 1
	<-ch
	wg.Done()
}

func CriticalSmutex() {
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrValue(&wg, ch)

	}
	wg.Wait()
	fmt.Println(x123)

}
