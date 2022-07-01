package main

import (
	"fmt"
)

func main() {

	test2()
}
func test2() {

	mymap := map[string]int{"b1": 22, "b2": 33, "b3": 44}
	fmt.Println(mymap)
	for key, val := range mymap {
		fmt.Printf("%v---%v\n", key, val)
	}
	delete(mymap,"b6")

}
func test1() {
	var a [5]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 5
	fmt.Println(a)
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("type of a:%T\n", a)
}
