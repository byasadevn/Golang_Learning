package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup
	fmt.Println("entry for main")
	wg.Add(1)
	go mygoroytine(ch, &wg)
	//wg.Wait()
	fmt.Println(<-ch)
	ch <- false
	wg.Wait()

}

func mygoroytine(a chan bool, wg *sync.WaitGroup) {
	fmt.Println("Entry point for mygoroytine")
	a <- true
	//close(a)
	//time.Sleep(10 * time.Second)
	fmt.Println("After chan write")
	fmt.Println("in my", <-a)
	wg.Done()
}
