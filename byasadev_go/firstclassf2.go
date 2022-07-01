package main

import "fmt"

type student struct {
	name  string
	class string
	grade string
}

func myfilter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func firstclassf2() {

	s1 := student{"byasa", "5th", "A+++"}
	s2 := student{"byasa1", "6th", "A+++"}
	s3 := student{"byasa2", "7th", "A++"}
	s4 := student{"byasa3", "7th", "A+++"}
	s5 := student{"byasa4", "7th", "A++"}

	students := []student{s1, s2, s3, s4, s5}

	basu := func(s student) bool {
		if s.grade == "A+++" {
			return true
		} else {
			return false
		}
	}

	fmt.Println(myfilter(students, basu))

	//f := myfilter(students, basu)

	//fmt.Println(f)

	/*f := myfilter(students, func(s student) bool {
		if s.class == "7th" {
			return true
		} else {
			return false
		}
	})*/

	//fmt.Println(f)

}
