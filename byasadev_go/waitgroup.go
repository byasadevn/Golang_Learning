package main

import (
	"fmt"
	"sync"
	"time"
)

func count(val int, wg *sync.WaitGroup) {

	fmt.Printf("The value received:%v\n", val)
	time.Sleep(1 * time.Second)
	fmt.Println("Wakeup from sleep", val)
	wg.Done()
}
func waitgroup() {
	fmt.Println("The main function")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go count(i, &wg)

	}
	wg.Wait()

	fmt.Println("The main is done")

}
