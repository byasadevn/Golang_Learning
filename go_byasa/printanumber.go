package main

import "fmt"

func main1() {

	for i := 33; i < 133; i++ {
		fmt.Printf("%#U\t", i)
	}
	var a interface{} = 23
	fmt.Println("the interface", a)

	x := 25
	if x == 23 {
		fmt.Println(x)
	} else if x = 24; x == 24 {
		fmt.Println("else part")
	}

	/*const (
		_ = iota
		x = 2018 + iota
		y = 2018 + iota
		z = 2018 + iota
		a = 2018 + iota

	)

	fmt.Println(x, y, z, a)
	*/
	/*qaz := `vgdasg
	ccdsdsconstdsvdvhvconstbvconst
	vcdsvvfconst\\new()

	dhgcdsgvconst`

	fmt.Println(qaz)
	*/
	/*x := 42
	fmt.Printf("%d\t%b\t%o\t%x\n", x, x, x, x)
	y := 42 >> 1
	fmt.Printf("%#d\t%#b\t%#o\t%#x\n", y, y, y, y)
	*/
	/*const a = 32
	const b int = 32
	fmt.Println(a, b)
	var c int = 12
	fmt.Println(c)
	*/
	/*a := (12 == 12)
	b := (12 <= 12)
	c := (12 >= 12)
	d := (12 != 12)
	e := (12 < 12)
	f := (12 > 12)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	*/
	//fmt.Printf("%v\n",a)

	//fmt.Printf( "basu:- %v\t%v\t%v\t%v\t%v\t%v\", a, b, c, d, e, f)
	//var x int = 123
	//fmt.Printf("%d\t%b\t%x", x, x, x)
}
