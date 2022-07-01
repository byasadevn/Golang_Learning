package main

import (
	"fmt"
)

type income interface {
	calculate() int
	source() string
}

type contractBilling struct {
	name         string
	cost         int
	pricePerHour int
}

type fixedBilling struct {
	name string
	cost int
}
type timeMaterial struct {
	name        string
	costperhour int
	total_hour  int
}

func (fb fixedBilling) calculate() int {
	return fb.cost
}
func (fb fixedBilling) source() string {
	return fb.name
}

func (tm timeMaterial) calculate() int {
	return tm.total_hour * tm.costperhour
}
func (tm timeMaterial) source() string {
	return tm.name
}

func (ct contractBilling) calculate() int {
	return ct.cost * ct.pricePerHour

}

func (ct contractBilling) source() string {
	return ct.name
}

func totalCalculation(in []income) {
	var totalv = 0
	for _, v := range in {
		totalv = totalv + v.calculate()
		fmt.Printf("The %v project has %v income\n", v.source(), v.calculate())
	}
	fmt.Println("The total imcome of the comaony:- ", totalv)
}

func interface4() {
	fb1 := fixedBilling{"Ivanti", 23}
	fb2 := fixedBilling{"CIvanti", 25}
	tm1 := timeMaterial{"IBM", 123, 5}
	tm2 := timeMaterial{"Watson", 34, 67}

	ct1 := contractBilling{"qzz1", 23, 23}
	ct2 := contractBilling{"wsx", 4, 5}

	in := []income{fb1, fb2, tm1, tm2, ct1, ct2}
	totalCalculation(in)

}
