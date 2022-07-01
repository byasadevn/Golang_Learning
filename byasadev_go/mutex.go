package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrX(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}

func mutex1() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go incrX(&wg, &mu)

	}
	wg.Wait()

	fmt.Println(x)

}
