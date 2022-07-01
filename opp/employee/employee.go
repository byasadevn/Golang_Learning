package employee

import (
	"fmt"
)

type Employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leaveTaken  int
}

func (e Employee) RemainingLeaves() {
	fmt.Printf("%s %s has %v remaining leaves", e.firstName, e.lastName, e.totalLeaves-e.leaveTaken)
}
