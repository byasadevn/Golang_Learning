package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func defer() {
	test2()
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Println("length is less than zero")
		return
	}
	if r.width < 0 {
		fmt.Println("width is less than zero")
		return
	}

	area := r.length * r.width
	fmt.Printf("The rect %v has the area: %d\n", r, area)

}

func test2() {
	r1 := rect{2, 3}
	r2 := rect{-3, 4}
	r3 := rect{4, 5}
	r4 := rect{5, -6}
	r5 := rect{-6, 7}

	rects := []rect{r1, r2, r3, r4, r5}
	var wg sync.WaitGroup
	for _, v := range rects {
		wg.Add(1)
		v.area(&wg)
	}
	wg.Wait()

}
func test1() {
	name := "BYASADEV"

	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
