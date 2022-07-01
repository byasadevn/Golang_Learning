package employee

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeaveTaken  int
}

func (e Employee) RemainingLeaves() {
	fmt.Printf("%s %s has %v remaining leaves", e.FirstName, e.LastName, e.TotalLeaves-e.LeaveTaken)
}
