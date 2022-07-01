package main

import (
	"fmt"
)

func display(i interface{}) {
	fmt.Printf("The interafce type %T and value %v\n", i, i)
}

func assert(i interface{}) {
	v, _ := i.(int)
	fmt.Println("The avlue of v and ok", v)
}

func interface3() {

	i1 := 24
	assert(i1)
	s1 := "Rakesh"
	assert(s1)

	/*var i int = 55
	display(i)

	var s string = "PSL"
	display(s)

	st := struct {
		name string
		id   int
	}{
		name: "persistent",
		id:   123,
	}
	display(st)*/

}
