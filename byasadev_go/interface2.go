package main

import (
	"fmt"
)

type salaryCalculator interface {
	calcualteSalary() int
}

type permanet struct {
	id     uint
	salary int
	pf     int
}

type contract struct {
	id     uint
	salary int
}

type freeLance struct {
	id          uint
	hour        int
	perhourrate int
}

func (a permanet) calcualteSalary() int {
	return a.salary + a.pf
}

func (c contract) calcualteSalary() int {
	return c.salary
}

func (f freeLance) calcualteSalary() int {
	return f.hour * f.perhourrate
}

func totalExpense(e []salaryCalculator) {
	total := 0
	for _, v := range e {
		total += v.calcualteSalary()
		/*fmt.Println(v.calcualteSalary())*/
	}
	fmt.Printf("The total expense: %v", total)

}

func interafce2() {

	p1 := permanet{1, 12, 2}
	p2 := permanet{2, 13, 3}
	c1 := contract{3, 20}
	c2 := contract{4, 30}

	f1 := freeLance{5, 2, 3}
	f2 := freeLance{6, 3, 4}

	emp := []salaryCalculator{p1, p2, c1, c2, f1, f2}
	totalExpense(emp)

}
