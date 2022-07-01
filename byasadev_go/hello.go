package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

func hello() {
	var c int64 = 5
	fmt.Printf("Sizeof=%d\n", unsafe.Sizeof(c))

	fmt.Println("------------------------------------")

	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	//defaultBool = customBool          //not allowed

	fmt.Println(trueConst, defaultBool, customBool)

	var qaz1, qaz2 = preremeter(2.3)
	fmt.Println(qaz1, qaz2)
	fmt.Println("----------------------------------------------")
	testSwitch(6)
	fmt.Println("----------------------------------------------")
	breakTheOuterLoop()

}

func preremeter(radius float64) (area, pere float64) {
	const pi float64 = 3.1412
	var wsx = 2 * 2
	area = pi * radius * radius

	pere = 2 * pi * radius * float64(wsx)
	return
}

func testSwitch(value int) {

	fmt.Println("The switch value =", value)
	switch value = 7; value {
	case 1:
		fmt.Println("Thump fringer: ", value)
	case 2:
		fmt.Println("Index fringer: ", value)
	case 3:
		fmt.Println("Middle fringer: ", value)
		fallthrough
	case 4:
		fmt.Println("Ring fringer: ", value)
	case 5:
		fmt.Println("Pinky fringer: ", value)
		fallthrough
	default:
		fmt.Println("Default fringer: ", value)

	}
	fmt.Println("The switch value end =", value)

}

func breakTheOuterLoop() {
myloop:
	for {
		//val := rand.Intn(100)
		//switch val{
		switch val := rand.Intn(100); {
		case val%2 == 0:
			fmt.Println("The val=", val)
			break myloop
		}
	}

}
